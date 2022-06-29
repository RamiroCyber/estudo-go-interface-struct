package cache

import "github.com/ramiro/generics/blog/entities"

type Cached interface {
	entities.Category | entities.Post
}

type cache[T Cached] struct {
	data map[string]T
}

func (c *cache[T]) Set(key string, value T) {
	c.data[key] = value
}

func (c *cache[T]) Get(key string) (v T) {
	if v, ok := c.data[key]; ok {
		return v
	}
	return
}

func New[T Cached]() *cache[T] {
	c := cache[T]{}
	c.data = make(map[string]T)
	return &c
}
