package pkg

type ResponseResult struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type LoggerResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
