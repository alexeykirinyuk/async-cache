package storage

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type CacheWithMetrics interface {
	Cache
	Metrics
}
