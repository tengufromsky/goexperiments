package performance

import "testing"

var benchmarkDynamicLen_GetLenMediumListenchVar int

func BenchmarkDynamicLen_GetLenMediumList(b *testing.B) {
	ll := dynamicLen{
		items: mediumList,
	}
	var res int
	for i := 0; i < b.N; i++ {
		res = ll.GetLen()
	}
	benchmarkDynamicLen_GetLenMediumListenchVar = res
}

var benchmarkLockedLen_GetLenMediumListVar int

func BenchmarkLockedLen_GetLenMediumList(b *testing.B) {
	ll := lockedLen{
		items:    mediumList,
		itemsLen: 10,
	}
	var res int
	for i := 0; i < b.N; i++ {
		res = ll.GetLen()
	}
	benchmarkLockedLen_GetLenMediumListVar = res
}

var benchmarkLockedLen_GetLenShortListVar int

func BenchmarkLockedLen_GetLenShortList(b *testing.B) {
	ll := lockedLen{
		items:    shortList,
		itemsLen: 5,
	}
	var res int
	for i := 0; i < b.N; i++ {
		res = ll.GetLen()
	}
	benchmarkLockedLen_GetLenShortListVar = res
}

var benchmarkDynamicLen_GetLenShortListVar int

func BenchmarkDynamicLen_GetLenShortList(b *testing.B) {
	ll := dynamicLen{
		items: shortList,
	}
	var res int
	for i := 0; i < b.N; i++ {
		res = ll.GetLen()
	}
	benchmarkDynamicLen_GetLenShortListVar = res
}

var benchmarkLockedLen_GetLenLongListVar int

func BenchmarkLockedLen_GetLenLongList(b *testing.B) {
	ll := lockedLen{
		items:    bigList,
		itemsLen: 100,
	}
	var res int
	for i := 0; i < b.N; i++ {
		res = ll.GetLen()
	}
	benchmarkLockedLen_GetLenLongListVar = res
}

var benchmarkDynamicLen_GetLenLongListVar int

func BenchmarkDynamicLen_GetLenLongList(b *testing.B) {
	ll := dynamicLen{
		items: bigList,
	}
	var res int
	for i := 0; i < b.N; i++ {
		res = ll.GetLen()
	}
	benchmarkDynamicLen_GetLenLongListVar = res
}

var shortList = []int{1, 2, 3, 4, 5}
var mediumList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var bigList []int

func init() {
	for i := 0; i < 100; i++ {
		bigList = append(bigList, i)
	}
}
