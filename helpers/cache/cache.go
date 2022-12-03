package cache

// Cache is a common-cache interface.
type Cache[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Has(key K) (ok bool)
	Set(key K, val V)
	Delete(key K)
	Values() []V
	Len() int
}

type SimpleCache[K comparable, V any] struct {
	items map[K]*item[V]
}

type item[V any] struct {
	value V
}

func New[K comparable, V any]() SimpleCache[K, V] {
	return SimpleCache[K, V]{
		items: make(map[K]*item[V], 0),
	}
}

func (c SimpleCache[K, V]) Set(key K, value V) {
	c.items[key] = &item[V]{value: value}
}

func (c SimpleCache[K, V]) Get(key K) (value V, ok bool) {
	item, ok := c.items[key]

	if !ok {
		return
	}

	return item.value, ok
}

func (c SimpleCache[K, V]) GetOrSet(key K, fallbackValue V) (value V, ok bool) {
	if val, ok := c.Get(key); ok {
		return val, ok
	}

	c.Set(key, fallbackValue)
	return c.Get(key)
}

func (c SimpleCache[K, V]) Has(key K) (ok bool) {
	_, ok = c.items[key]
	return ok
}

func (c SimpleCache[K, V]) Delete(key K) {
	delete(c.items, key)
}

func (c SimpleCache[K, V]) Len() int {
	return len(c.items)
}

func (c SimpleCache[K, V]) Keys() []K {
	result := make([]K, 0, len(c.items))
	for key, _ := range c.items {
		result = append(result, key)
	}
	return result
}

func (c SimpleCache[K, V]) Values() []V {
	result := make([]V, 0, len(c.items))
	for _, value := range c.items {
		result = append(result, value.value)
	}
	return result
}
