package helpers

type GeneralResponse struct {
	Status  bool        `json:"status" form:"status"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
