package alg

// Union-Find 算法是在 Weighted Quick-Union 的基础之上进一步优化，即路径压缩的 Weighted Quick-Union 算法
// 所谓 路径压缩 ，就是在计算结点 i 的根之后，将回溯路径上的每个被检查结点的 id 设置为root(i)

func PathCompressionRoot(i int32) {
	root := root(i)
	for i != root {
		t := a[i]
		a[i] = root
		i = t
	}
}

// InsertSort 插入排序(增序)
func InsertSort(data []int32) {
	n := len(data)
	for i := 1; i < n; i++ {
		v := data[i]
		j := i - 1
		for j > 0 && v < data[j] {
			data[j+1] = data[j]
			j = j - 1
		}

		data[j+1] = v
	}
}

// MergeSort 合并排序
func MergeSort(data []int32, s, e int) {
	if s == e {
		return
	}

	m := (s + e) / 2
	MergeSort(data, s, m)
	MergeSort(data, m+1, e)

	Merge(data, s, m, e)
}

// Merge 合并两个已经排序的数组
func Merge(data []int32, s, m, e int) {
	// 左指针
	p1 := s
	// 右指针
	p2 := m + 1

	for p1 != p2 {
		// 左边数据大于右边数据，将右边数据插入
		if data[p1] > data[p2] {
			v := data[p2]
			for i := p2; i > p1; i-- {
				data[i] = data[i-1]
			}
			data[p1] = v
			if p2 < e {
				p2 = p2 + 1
			}
		}
		p1 = p1 + 1
	}
}

// 斐波那契数列
func Fib(n uint64) uint64 {
	if n < 2 {
		return n
	}

	first := uint64(0)
	second := uint64(1)
	third := uint64(0)
	for i := uint64(2); i <= n; i++ {
		third = first + second
		first = second
		second = third
	}

	return third
}
