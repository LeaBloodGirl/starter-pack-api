package cache

import (
	"encoding/json"
	"net/http"
	"starter-pack-api/internal/config"
	"starter-pack-api/internal/logger"
	"time"

	"github.com/patrickmn/go-cache"
)

type AllCache struct {
	memoryCache *cache.Cache
}

func NewCache(c config.CacheParams) *AllCache {
	cache := cache.New(time.Duration(c.ExpirationTime)*time.Second, time.Duration(c.ExpurgeTime)*time.Second)
	/*cachesize := 1024 * 1024 * 1024
	cache := freecache.NewCache(cachesize)*/
	return &AllCache{
		memoryCache: cache,
	}
}

func (c *AllCache) Read(id string) (item []byte, err error, ok bool) {
	cacheGetted, ok := c.memoryCache.Get(id)
	if ok {
		res, err := json.Marshal(cacheGetted)
		if err != nil {
			return nil, err, false
		}
		return res, nil, true
	}
	return nil, nil, false
}

func (c *AllCache) Update(id string, element interface{}, cacheDuration time.Duration) {
	c.memoryCache.Set(id, element, cacheDuration)
}

func (c *AllCache) CheckCache(logDebug logger.Logger, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err, ok := c.Read(r.RequestURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logDebug.ErrorLevel("Error while reading cache : ", err)
			return
		}
		if ok {
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}
		f(w, r)
	}
}
