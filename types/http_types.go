package types

type HeaderGetter interface {
	Get(key string, defaultValue ...string) string
}
