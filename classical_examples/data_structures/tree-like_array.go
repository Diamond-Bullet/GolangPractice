package data_structures

// reference: https://blog.csdn.net/bestsort/article/details/80796531

type TreeLikeArray struct {
	N                        int
	RawArray, ProcessedArray []int
}

// LowBit in the complement mode, (-x)'s binary expression is like ^(-x)+1
// eg: 10: 0b00001010 -10: 0b10001010 0b11110101 0b11110110 (use biu.ToBinaryString to print it)
// reference: https://blog.csdn.net/zl10086111/article/details/80907428/
func LowBit(x int) int {
	return x & (-x)
}

func NewTreeLikeArray(num []int) *TreeLikeArray {
	t := &TreeLikeArray{
		N:              len(num),
		RawArray:       num,
		ProcessedArray: make([]int, len(num)+1),
	}

	pre := make([]int, t.N+1)
	for i := 1; i <= t.N; i++ {
		pre[i] = pre[i-1] + num[i]
		t.ProcessedArray[i] = pre[i] - pre[i-LowBit(i)]
	}

	return t
}

func (t *TreeLikeArray) Update(i, k int) {
	delta := k - t.RawArray[i-1]
	for j := i; i <= t.N; {
		t.ProcessedArray[j] += delta
		j += LowBit(j)
	}
	t.RawArray[i-1] = k
}

func (t *TreeLikeArray) GetSum(i int) int {
	var res int
	for i > 0 {
		res += t.ProcessedArray[i]
		i -= LowBit(i)
	}
	return res
}
