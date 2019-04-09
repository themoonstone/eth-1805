package _1_testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(10,20)
	t.Log(res)
	if res != 20 {
		// 中断后续执行
		t.Fatalf("add num 10 and 20 error")
	}
	t.Logf("test add successed!")
}
