package problem0047

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]int
}

func Test_Problem0047(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{1, 2, 3}},
			ans{[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			}},
		},

		{
			para{[]int{1, 1, 2}},
			ans{[][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			}},
		},

		{
			para{[]int{5, 5, 5, 5, 5}},
			ans{[][]int{
				{5, 5, 5, 5, 5},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, permuteUnique(p.one), "输入:%v", p)
	}
}
