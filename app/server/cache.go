package server

import (
	"errors"
	"github.com/ReneKroon/ttlcache"
)

var Cache *ttlcache.Cache

func InitCache() {
	Cache = ttlcache.NewCache()
}

func GetCache() (*ttlcache.Cache, error) {
	if Cache != nil {
		return Cache, nil
	}
	return Cache, errors.New("cache is nil")
}
