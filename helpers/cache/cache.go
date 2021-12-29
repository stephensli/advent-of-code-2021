package cache

// Cache is a common-cache interface.
type Cache[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Has(key K) (ok bool)
	Set(key K, val V)
	Delete(key K)
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

func (c SimpleCache[K, V]) Has(key K) (ok bool) {
	_, ok = c.items[key]
	return ok
}

func (c SimpleCache[K, V]) Delete(key K) {
	delete(c.items, key)
}
