package cache

import (
	"blog/database"
	"github.com/vmihailenco/msgpack/v4"

	"github.com/go-redis/cache/v7"
)

var cached *cache.Codec

func GetCache() *cache.Codec {
	if cached == nil {
		cached = initCache()
	}
	return cached
}

// 初始化缓存模块 并注入缓存驱动
func initCache() *cache.Codec {
	return &cache.Codec{
		Redis: database.GetRedis(),

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
}
