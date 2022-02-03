package sache

import "github.com/satheesh1997/sache/policies"

type (
	Sache struct {
		EvictionPolicy     policies.LRUEvictionPolicy
		Storage            map[string]string
		StorageSize        int
		CurrentStorageSize int
	}
)

func (cache *Sache) Put(key string, value string) {
	if cache.CurrentStorageSize == cache.StorageSize {
		keyToRemove := cache.EvictionPolicy.EvictKey()
		delete(cache.Storage, keyToRemove)
	}

	cache.EvictionPolicy.KeyAccessed(key)
	cache.Storage[key] = value

	if cache.CurrentStorageSize < cache.StorageSize {
		cache.CurrentStorageSize++
	}
}

func (cache *Sache) Get(key string) string {
	value, found := cache.Storage[key]
	if !found {
		return ""
	}

	cache.EvictionPolicy.KeyAccessed(key)
	return value
}

func New(storageSize int) Sache {
	cache := Sache{
		EvictionPolicy:     policies.NewLRUEvictionPolicy(),
		Storage:            make(map[string]string),
		StorageSize:        storageSize,
		CurrentStorageSize: 0,
	}
	return cache
}
