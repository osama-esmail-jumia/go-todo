package context

type Context interface {
	BindQuery(obj interface{}) error
	Bind(obj interface{}) error
	Param(key string) string
}
