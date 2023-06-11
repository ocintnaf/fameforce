package types

type HeaderGetter interface {
	Get(key string, defaultValue ...string) string
}

type CtxLocalsGetter interface {
	Locals(key interface{}, value ...interface{}) interface{}
}
