package performance

import "testing"

func BenchmarkInterfaceCall(b *testing.B) {
	own := OwnAboveInterface{
		inner: &SimpleStruct{
			elements: bigList,
		},
	}
	for i := 0; i < b.N; i++ {
		res := own.Own()
		_ = res
	}
}

func BenchmarkStructCall(b *testing.B) {
	own := OwnAboveStruct{
		inner: &SimpleStruct{
			elements: bigList,
		},
	}
	for i := 0; i < b.N; i++ {
		res := own.Own()
		_ = res
	}
}
