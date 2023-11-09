package abstraction

type BlogCache interface {
	Get(key string) (string, error)
	Set(key string, val string) error
	Del(key string) error
	Exists(key string) bool
	Keys(prefixKey string) ([]string, error)
}
