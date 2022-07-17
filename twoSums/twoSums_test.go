package twosums

import "testing"

/*
goos: windows
goarch: amd64
pkg: github.com/arkiant/leetcode/twoSums
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkTwoSum-8   	33032188	        34.92 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/arkiant/leetcode/twoSums	1.407s
*/
func BenchmarkTwoSum(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum([]int{3, 2, 4}, 6)
	}
}
