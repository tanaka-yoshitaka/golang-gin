package controllers

type Handler struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewHandler(status int, message string, result interface{}) *Handler {
	H := new(Handler)
	H.Status = status
	H.Message = message
	H.Result = result
	return H
}
