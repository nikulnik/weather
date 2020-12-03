package cache

import (
	"testing"
	"time"

	"github.com/nikulnik/weather/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCache_ReturnsCache(t *testing.T) {
	t.Parallel()
	cache := NewCache(time.Second)
	assert.NotNil(t, cache)
}

func TestSetAndGet_GetsCorrectValue(t *testing.T) {
	t.Parallel()
	cache := NewCache(time.Minute)
	cache.Set("key", "val")
	assert.Equal(t, "val", cache.Get("key"))
}

func TestGet_ReturnsNilAfterTimeout(t *testing.T) {
	t.Parallel()
	cache := NewCache(time.Millisecond * 3)
	value := &domain.Forecast{Sunset: 41241}
	cache.Set("key", value)
	time.Sleep(time.Millisecond * 6)
	assert.Nil(t, cache.Get("key"))
}

func TestGet_ResetsTimeout(t *testing.T) {
	t.Parallel()
	cache := NewCache(time.Millisecond * 10)
	value := "val"
	cache.Set("key", value)
	time.Sleep(time.Millisecond * 5)
	cache.Get("key")
	time.Sleep(time.Millisecond * 5)
	assert.Equal(t, value, cache.Get("key"))
}
