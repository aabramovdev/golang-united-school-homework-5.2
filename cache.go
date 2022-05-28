package cache

import "time"

const (
	hoursInHundredYears = 876600
)

type Value struct {
	value          string
	expirationTime time.Time
}

type Cache struct {
	cache map[string]Value
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]Value),
	}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.cache[key]
	if time.Now().After(value.expirationTime) {
		delete(c.cache, key)
		return "", false
	}
	return value.value, ok
}

func (c Cache) Put(key, value string) {
	c.cache[key] = Value{value: value, expirationTime: time.Now().Add(time.Hour * hoursInHundredYears)}
}

func (c Cache) Keys() []string {
	var keys []string
	for key, value := range c.cache {
		if time.Now().Before(value.expirationTime) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadLine time.Time) {
	c.cache[key] = Value{value: value, expirationTime: deadLine}
}
