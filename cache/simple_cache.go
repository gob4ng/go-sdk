package cache

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type SimpleCacheSetup struct {
	PrefixKey   string
	ExpiredTime int
	PurgeTime   int
}

type simpleCacheConfig struct {
	simpleCacheSetup SimpleCacheSetup
	cache            cache.Cache
}

func NewSimpleCacheConfig(setup SimpleCacheSetup) (*simpleCacheConfig, *error) {

	newCache := cache.New(time.Millisecond*time.Duration(setup.ExpiredTime),
		time.Millisecond*time.Duration(setup.PurgeTime))

	config := simpleCacheConfig{
		simpleCacheSetup: setup,
		cache:            *newCache,
	}

	return &config, nil
}

func (c simpleCacheConfig) Set(key string, value string) {

	cacheKey := c.simpleCacheSetup.PrefixKey + "_" + key
	c.cache.Set(cacheKey, value, time.Millisecond*time.Duration(c.simpleCacheSetup.ExpiredTime))
}

func (c simpleCacheConfig) Get(key string) (*string, *error) {

	cacheKey := c.simpleCacheSetup.PrefixKey + "_" + key
	value, found := c.cache.Get(cacheKey)
	if found {
		strValue := fmt.Sprintf("%v", value)
		return &strValue, nil
	}

	newError := errors.New("cache not found")

	return nil, &newError
}

func (c simpleCacheConfig) Delete(key string) {
	cacheKey := c.simpleCacheSetup.PrefixKey + "_" + key
	c.cache.Delete(cacheKey)
}
