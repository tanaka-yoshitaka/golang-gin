package controllers

type Handler struct {
    Message string   `json:"message"`
    Data interface{} `json:"data"`
}

func NewHandler(message string, data interface{}) *Handler {
    H := new(Handler)
    H.Message = message
    H.Data = data
    return H
}