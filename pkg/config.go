package pkg

type Config struct {
	Port          string `json:"port"`
	AccessLogPath string `json:"access_log_path"`
	LogLevel      string `json:"log_level"`
	Env           string `json:"env"`
	Pgsql         PgSQL  `json:"pgsql"`
}

type PgSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}
