package performance

import "testing"

func BenchmarkDynamicLen_GetLenMediumList(b *testing.B) {
	ll := dynamicLen{
		items: mediumList,
	}
	for i := 0; i < b.N; i++ {
		res := ll.GetLen()
		_ = res
	}
}

func BenchmarkLockedLen_GetLenMediumList(b *testing.B) {
	ll := lockedLen{
		items:    mediumList,
		itemsLen: 10,
	}
	for i := 0; i < b.N; i++ {
		res := ll.GetLen()
		_ = res
	}
}

func BenchmarkLockedLen_GetLenShortList(b *testing.B) {
	ll := lockedLen{
		items:    shortList,
		itemsLen: 5,
	}
	for i := 0; i < b.N; i++ {
		res := ll.GetLen()
		_ = res
	}
}

func BenchmarkDynamicLen_GetLenShortList(b *testing.B) {
	ll := dynamicLen{
		items: shortList,
	}
	for i := 0; i < b.N; i++ {
		res := ll.GetLen()
		_ = res
	}
}

func BenchmarkLockedLen_GetLenLongList(b *testing.B) {
	ll := lockedLen{
		items:    bigList,
		itemsLen: 100,
	}
	for i := 0; i < b.N; i++ {
		res := ll.GetLen()
		_ = res
	}
}

func BenchmarkDynamicLen_GetLenLongList(b *testing.B) {
	ll := dynamicLen{
		items: bigList,
	}
	for i := 0; i < b.N; i++ {
		res := ll.GetLen()
		_ = res
	}
}

var shortList = []int{1, 2, 3, 4, 5}
var mediumList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var bigList []int

func init() {
	for i := 0; i < 100; i++ {
		bigList = append(bigList, i)
	}
}
