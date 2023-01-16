package questions

import (
	"math"
	"math/rand"
)

// CountingSort MarkIt 计数排序
func CountingSort(arr []int, k int, locate func(num int) int) []int {
	rank := make([]int, k)
	for _, x := range arr {
		rank[locate(x)] += 1
	}
	for i := 1; i < k; i++ {
		rank[i] += rank[i-1]
	}

	n := len(arr)
	ret := make([]int, n)
	var location int
	for i := n - 1; i >= 0; i-- {
		location = locate(arr[i])
		ret[rank[location]-1] = arr[i]
		rank[location]--
	}
	return ret
}

// RadixSort MarkIt 基数排序
func RadixSort(arr []int, base, d int) []int {
	keys := map[int][]int{}
	for _, x := range arr {
		keys[x] = []int{x, -1}
	}
	for i := 0; i < d; i++ {
		arr = CountingSort(arr, base, func(num int) int {
			if keys[num][1] == -1 {
				divisor := int(math.Pow(float64(base), float64(i+1)))
				keys[num][0], keys[num][1] = keys[num][0]/divisor, keys[num][0]%divisor
				return keys[num][1]
			}
			ret := keys[num][1]
			keys[num][1] = -1
			return ret
		})
	}
	return arr
}

// QuickSort MarkIt 快速排序
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	boundary := Partition(arr)
	QuickSort(arr[:boundary])
	QuickSort(arr[boundary+1:])
}

func Partition(arr []int) int {
	pivot := arr[0]
	j, n := 0, len(arr)
	for i := 1; i < n; i++ {
		if arr[i] <= pivot {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[0], arr[j] = arr[j], arr[0]
	return j
}

// HeapSort MarkIt 堆排序
func HeapSort(arr []int) {
	Heapify(arr)

	n := len(arr)
	for n > 0 {
		arr[0], arr[n-1] = arr[n-1], arr[0]
		arr = arr[:n-1]
		SiftDown(arr, 0)
		n--
	}
}

// Heapify 堆初始化
func Heapify(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		SiftDown(arr, i)
	}
}

// SiftDown 堆结点移动
func SiftDown(arr []int, i int) {
	n := len(arr)
	for i <= n/2-1 {
		changeChild := 2*i + 1
		if 2*i+2 < n && arr[2*i+2] < arr[2*i+1] {
			changeChild = 2*i + 2
		}
		if arr[i] <= arr[changeChild] {
			break
		}
		arr[i], arr[changeChild] = arr[changeChild], arr[i]
		i = changeChild
	}
}

// FindNumK MarkIt 寻找第k小的数，期望O(n)，最坏O(n)
func FindNumK(arr []int, l, r, k int) int {
	return findNumK(arr, l, r, k-1)
}

func findNumK(arr []int, l, r, k int) int {
	if r-l <= 5 {
		InsertSort(arr, l, r)
		return arr[l+k]
	}
	group := (r - l) / 5
	for i := 0; i < group; i++ {
		left := l + i*5
		right := l + (i+1)*5
		if right > r {
			right = r
		}
		InsertSort(arr, left, right)
		arr[l+i], arr[(left+right)/2] = arr[(left+right)/2], arr[l+i]
	}
	x := findNumK(arr, l, l+group, l+group/2)
	divideLine := PartitionForFindNumK(arr, l, r, x)
	if divideLine == k {
		return arr[divideLine]
	} else if divideLine > k {
		return findNumK(arr, l, divideLine, k)
	}
	return findNumK(arr, divideLine, r, k-divideLine)
}

// InsertSort MarkIt 插入排序
func InsertSort(arr []int, l, r int) {
	for i := l + 1; i < r; i++ {
		key := arr[i]
		j := i - 1
		for ; j >= l; j-- {
			if arr[j] <= key {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func PartitionForFindNumK(arr []int, l, r, pivot int) int {
	for k := l; k < r; k++ {
		if arr[k] == pivot {
			pivot = k
			break
		}
	}
	arr[l], arr[pivot] = arr[pivot], arr[l]
	pivot = arr[l]
	i := l
	for j := l + 1; j < r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[i] = arr[i], arr[l]
	return i
}

// PartitionRepetitive MarkIt 针对重复元素的快速排序partition
func PartitionRepetitive(nums []int, p, r int) (int, int) {
	q, t := p, p
	pivot := nums[p]
	for i := p + 1; i < r; i++ {
		if nums[i] < pivot {
			nums[q], nums[i] = nums[i], nums[q]
			nums[t], nums[i] = nums[i], nums[t]
			t++
			q++
		} else if nums[i] == pivot {
			nums[t], nums[i] = nums[i], nums[t]
			t++
		}
	}
	nums[p], nums[t] = nums[t], nums[p]
	return q, t + 1
}

// LocalSort MarkIt 对有k个关键字的n条数据，线性空间原址排序
func LocalSort(arr [][2]int, k int) {
	sortBy := make([]int, k)
	for _, x := range arr {
		sortBy[x[1]] += 1
	}
	for i := 1; i < k; i++ {
		sortBy[i] += sortBy[i-1]
	}
	for j := len(arr) - 1; j >= 0; j-- {
		p := sortBy[arr[j][1]] - 1
		for p > j {
			sortBy[arr[j][1]] -= 1
			arr[p], arr[j] = arr[j], arr[p]
			p = sortBy[arr[j][1]] - 1
		}
	}
}

// FindNumKByQuickSort MarkIt 寻找第k小的数, 期望O(n)，最坏O(n**2)
func FindNumKByQuickSort(arr []int, l, r, k int) int {
	return findNumKByQuickSort(arr, l, r, k-1)
}

func findNumKByQuickSort(arr []int, l, r, k int) int {
	dividingLine := RandomPartition(arr, l, r)
	if dividingLine == k {
		return arr[dividingLine]
	} else if dividingLine > k {
		return findNumKByQuickSort(arr, l, dividingLine, k)
	}
	return findNumKByQuickSort(arr, dividingLine, r, k)
}

func RandomPartition(arr []int, l, r int) int {
	pivot := rand.Intn(r-l) + l
	arr[l], arr[pivot] = arr[pivot], arr[l]
	pivot = arr[l]
	i := l
	for j := l + 1; j < r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[i] = arr[i], arr[l]
	return i
}
