package problem2568

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	{
		[][]int{
			{0,2,1,0},
			{4,0,0,3},
			{1,0,0,4},
			{0,3,2,0},

		},
		7,
	},

	// 可以有多个 testcase
}

func Test_findMaxFish(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMaxFish(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_findMaxFish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMaxFish(tc.grid)
		}
	}
}
