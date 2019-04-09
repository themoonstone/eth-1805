package _3_bench_testing

import "testing"

func TestGetCountsOfPrimeNumber(t *testing.T) {
	t.Logf("primeNumber counts : %v\n", GetCountsOfPrimeNumber(100))
}

func TestGetCountsOfPrimeNumberWithSqrt(t *testing.T) {
	t.Logf("primeNumber counts : %v\n", GetCountsOfPrimeNumberWithSqrt(100))
}

// 压力测试
func BenchmarkGetCountsOfPrimeNumber(b *testing.B) {
	// 汇报内存的使用情况
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetCountsOfPrimeNumber(1000)
	}
	// 输出结果:
	//										函数运行次数			  单次操作花费的时间(ns)
	// BenchmarkGetCountsOfPrimeNumber-4        2000000               684 ns/op

}

func BenchmarkGetCountsOfPrimeNumberWithSqrt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetCountsOfPrimeNumberWithSqrt(1000)
	}
}

/*
	// 数量级为100
	BenchmarkGetCountsOfPrimeNumber-4                 200000             10588 ns/op
	BenchmarkGetCountsOfPrimeNumberWithSqrt-4        1000000              2144 ns/op

	// 数量级为200
	BenchmarkGetCountsOfPrimeNumber-4                  30000             41125 ns/op
	BenchmarkGetCountsOfPrimeNumberWithSqrt-4         300000              5627 ns/op

*/