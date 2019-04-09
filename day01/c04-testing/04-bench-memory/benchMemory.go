package _4_bench_memory

func AllocMemory() []byte {
	return make([]byte, 1024*10)
}