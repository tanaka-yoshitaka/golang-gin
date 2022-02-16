package controllers

// interfaces層で使うgin.Contextのメソッドを追加する
type Context interface {
	AbortWithStatus(code int)
	Get(key string) (value interface{}, exists bool)
	JSON(code int, obj interface{})
	Param(key string) string
	PostForm(key string) string
	MustGet(key string) interface{}
	Bind(obj interface{}) error
}
