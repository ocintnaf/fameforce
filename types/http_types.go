package types

type HeaderGetter interface {
	Get(key string, defaultValue ...string) string
}

type CtxLocaler interface {
	Locals(key interface{}, value ...interface{}) interface{}
}
