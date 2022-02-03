package sache

import "github.com/satheesh1997/sache/policies"

type (
	Sache struct {
		EvictionPolicy policies.LRUEvictionPolicy
		Storage        HashMap
	}
)

func (cache *Sache) Put(key string, value string) {
	if cache.Storage.IsFull() {
		keyToRemove := cache.EvictionPolicy.EvictKey()
		cache.Storage.Remove(keyToRemove)
	}
	cache.EvictionPolicy.KeyAccessed(key)
	cache.Storage.Insert(key, value)
}

func (cache *Sache) Get(key string) string {
	value := cache.Storage.Read(key)
	if value != "" {
		cache.EvictionPolicy.KeyAccessed(key)
	}
	return value
}

func New(storageSize int) Sache {
	cache := Sache{
		EvictionPolicy: policies.NewLRUEvictionPolicy(),
		Storage: HashMap{
			Data:    make(map[string]string),
			Size:    0,
			MaxSize: storageSize,
		},
	}
	return cache
}
