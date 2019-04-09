package _3_bench_testing

import "math"

// 获取指定范围的素数的数量
func GetCountsOfPrimeNumber(n int) int {
	var counts int = 0
	// 最小的素数是2。0和1不是素数
	//
	for i := 2; i <= n; i++  {
		var isPrimeNumber bool
		for j := 2; j < i; j++ {
			if i % j == 0 {
				isPrimeNumber = true
				break
			}
		}
		if !isPrimeNumber {
			counts++
		}
	}
	return counts
}

// 优化算法
// 判断一个数是否是素数，判断从0到当前数值开方之间是否有能够被它整除的数
// 判断17是否是素数
// 1. 开方 得到4
// 2. 如果说17不能被4以下的数整除，就说明它是一个素数
// TODO 自己理解上述结论的产生原因
func GetCountsOfPrimeNumberWithSqrt(n int) int {
	var counts int = 0
	// 最小的素数是2。0和1不是素数
	//
	var k, j int
	for i := 2; i <= n; i++  {
		// 得到平方根
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i % j == 0 {
				break
			}
		}
		// 满足素数条件
		if j > k {
			counts++
		}
	}
	return counts
}