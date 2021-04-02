package alg

import "fmt"

// 并查集支持如下操作：
// 查询：查询某个元素属于哪个集合，通常是返回集合内的一个“代表元素”。这个操作是为了判断两个元素是否在同一个集合之中。
// 合并：将两个集合合并为一个。
// 对象的不相交集（Disjoint Set）
// Find 查询：两个对象是否在同一集合中？
// Union 命令：将包含两个对象的集合替换为它们的并集。

var (
	a = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func Find(p, q int32) bool {
	return a[p] == a[q]
}

func Union(p, q int32) {
	pid := a[p]
	for i := 0; i < len(a); i++ {
		if a[i] == pid {
			a[i] = a[q]
		}
	}
}

func Print() {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}
