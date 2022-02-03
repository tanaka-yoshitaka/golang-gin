package controllers

type Context interface {
	Param(key string) string
	PostForm(key string) string
	JSON(code int, obj interface{})
	AbortWithStatus(code int)
	MustGet(key string) interface{}
}
