package main

import (
"fmt"
"sync"
)

type Cache interface {
    Set(key string, value string)
    Get(key string) (string, bool)
    Delete(key string)
}

type myCache struct {
    mu    sync.Mutex
    cache map[string]string
}

func NewCache() Cache {
    return &myCache{
        cache: make(map[string]string),
    }
}

func (c *myCache) Set(key string, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = value
}

func (c *myCache) Get(key string) (string, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    value, ok := c.cache[key]
    return value, ok
}

func (c *myCache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    delete(c.cache, key)
}

func main() {
    cache := NewCache()
    cache.Set("key1", "value1")
    cache.Set("key2", "value2")

    val1, ok := cache.Get("key1")
    if ok {
	    fmt.Println("Value for key1:", val1)
    }

    val2, ok := cache.Get("key2")
    if ok {
	    fmt.Println("Value for key2:", val2)
    }

    cache.Delete("key1")

    val1, ok = cache.Get("key1")
    if ok {
	    fmt.Println("Value for key1 after deletion:", val1)
    } else {
	    fmt.Println("Key1 not found in cache")
    }
}