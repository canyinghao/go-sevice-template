# go-gin服务端模板

### 技术栈

go + gin + pgsql + zap


### 安装依赖
执行make install 安装依赖包

### 本地运行
执行make local.run

###  swagger文档
在电脑上安装swag命令
go install github.com/swaggo/swag/cmd/swag@latest

修改注释后，需要用swag init命令，重新生成文档


### grpc支持
rpc需要在config中配置一个端口才会启用  
gin使用一个端口，rpc使用另一个端口  
