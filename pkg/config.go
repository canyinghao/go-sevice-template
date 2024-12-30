package pkg

type Config struct {
	Port          string                 `json:"port"`
	AccessLogPath string                 `json:"access_log_path"`
	Rpc           string                 `json:"rpc"`
	Cron          bool                   `json:"cron"`
	LogLevel      string                 `json:"log_level"`
	Env           string                 `json:"env"`
	Pgsql         DbConfig               `json:"pgsql"`
	Redis         map[string]RedisConfig `json:"redis"`
}

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
