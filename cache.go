package main

import (
	"net/http"
	"sync"
)

type CacheEntry struct {
	statusCode int
	body       []byte
	headers    http.Header
}

var (
	cache = make(map[string]CacheEntry)
	mutex sync.RWMutex
)

func addToCache(key string, entry CacheEntry) {
	mutex.Lock()
	defer mutex.Unlock()
	cache[key] = entry
}

func getFromCache(path string) (CacheEntry, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	entry, ok := cache[path]
	return entry, ok
}

func sendCachedResponse(w http.ResponseWriter, entry CacheEntry) {
	for key, values := range entry.headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.Header().Add("X-Pranay-Cache", "HIT")
	w.WriteHeader(entry.statusCode)
	w.Write(entry.body)
}
