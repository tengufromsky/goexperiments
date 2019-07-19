package performance

import (
	"context"
	"sync"
)

type OwnAboveInterface struct {
	inner SimpleInterface
}

func (o *OwnAboveInterface) Own() int {
	return o.inner.Method(context.Background())
}

type SimpleInterface interface {
	Method(ctx context.Context) int
}

type OwnAboveStruct struct {
	inner *SimpleStruct
}

func (o *OwnAboveStruct) Own() int {
	return o.inner.Method(context.Background())
}

type SimpleStruct struct {
	sync.Mutex
	c        int
	elements []int
}

func (s *SimpleStruct) Method(ctx context.Context) int {
	s.Lock()
	defer s.Unlock()
	c := s.c % 100
	return s.elements[c]
}
