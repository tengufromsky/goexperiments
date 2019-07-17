package performance

import "sync/atomic"

type lockedLen struct {
	items    []int
	itemsLen int
	counter  uint64
}

func (r *lockedLen) GetLen() int {
	return int(atomic.AddUint64(&r.counter, 1) % uint64(r.itemsLen))
}

type dynamicLen struct {
	items   []int
	counter uint64
}

func (r *dynamicLen) GetLen() int {
	return int(atomic.AddUint64(&r.counter, 1) % uint64(len(r.items)))
}
