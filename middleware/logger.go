package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog(c *pkg.Config) {
	// 1、encoder
	encndercfg := zap.NewProductionEncoderConfig()
	encndercfg.TimeKey = "time"                          // 改变时间的key
	encndercfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 更改时间格式
	encndercfg.EncodeLevel = zapcore.CapitalLevelEncoder //将日志级别大写并带有颜色
	enconder := zapcore.NewJSONEncoder(encndercfg)
	// 2、writerSyncer 将日志写到文件和终端
	file, _ := os.OpenFile(c.AccessLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)

	// 3、设置loglevel

	level := zapcore.DebugLevel

	if c.LogLevel == "debug" {
		level = zapcore.DebugLevel
	} else if c.LogLevel == "info" {
		level = zapcore.InfoLevel
	} else if c.LogLevel == "warn" {
		level = zapcore.WarnLevel
	} else if c.LogLevel == "error" {
		level = zapcore.ErrorLevel
	}

	// 创建zapcore
	core := zapcore.NewCore(enconder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// 创建logger
	logger := zap.New(core)

	// 替换zap全局的logger
	zap.ReplaceGlobals(logger)
	zap.L().Info("logger init success")

}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		body := c.Request.Header.Get("C_body")
		c.Request.Header.Del("C_body")

		result := c.Request.Header.Get("C_result")
		c.Request.Header.Del("C_result")

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("referer", c.Request.Header.Get("Referer")),
			zap.String("body", body),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
			zap.String("result", result),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
