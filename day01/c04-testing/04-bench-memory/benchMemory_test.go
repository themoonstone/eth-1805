package _4_bench_memory

import "testing"

func BenchmarkAllocMemory(b *testing.B) {
	//b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = AllocMemory()
	}
}
// go test -bench . -benchmem -gcflags "-N -l"
// -gcflags "-N -l":禁用内联优化
// BenchmarkAllocMemory-4           1000000
// 1117 ns/op
// 10240 B/op    (单次操作执行时，堆内存分配的问题)
// 1 allocs/op  （单次操作执行的时候，内存分配的次数）
