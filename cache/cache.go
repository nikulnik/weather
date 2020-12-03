package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type cache struct {
	data  map[string]value
	mux          *sync.Mutex
	ttl          time.Duration
}

type value struct {
	timer *time.Timer
	data  interface{}
}

func NewCache(ttl time.Duration) Cache {
	return &cache{
		data:  		make(map[string]value),
		mux:          &sync.Mutex{},
		ttl:          ttl,
	}
}

func (c *cache) Set(key string, val interface{}) {
	c.mux.Lock()
	timer := time.NewTimer(c.ttl)
	c.data[key] = value{
		data:  val,
		timer: timer,
	}
	c.mux.Unlock()
	go func() {
		<-timer.C
		c.mux.Lock()
		delete(c.data, key)
		c.mux.Unlock()
	}()
}

func (c *cache) Get(key string) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.data[key]; ok {
		c.data[key].timer.Reset(c.ttl)
	}
	return c.data[key].data
}
