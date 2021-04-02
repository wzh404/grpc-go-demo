package alg

import (
	"fmt"
	"testing"
)

func TestQuickFind(t *testing.T) {
	Union(3, 4)
	Union(4, 9)
	Union(8, 0)
	Union(2, 3)
	Union(5, 6)
	Union(5, 9)
	Union(7, 3)
	Union(4, 8)
	Union(6, 1)
	Print()
}

func TestQuickUnion(t *testing.T) {
	QUnion(3, 4)
	QUnion(4, 9)
	QUnion(8, 0)
	QUnion(2, 3)
	QUnion(5, 6)
	QUnion(5, 9)
	QUnion(7, 3)
	QUnion(4, 8)
	QUnion(6, 1)
	Print()
}

func TestWeightQuickUnion(t *testing.T) {
	WQUnion(3, 4)
	WQUnion(4, 9)
	WQUnion(8, 0)
	WQUnion(2, 3)
	WQUnion(5, 6)
	WQUnion(5, 9)
	WQUnion(7, 3)
	WQUnion(4, 8)
	WQUnion(6, 1)
	Print()
}

func TestMergeSort(t *testing.T) {
	a := []int32{1, 3, 5, 2, 6, 7, 4}
	MergeSort(a, 0, 6)
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func TestMerge(t *testing.T) {
	var d = []int32{1, 3, 2}
	Merge(d, 0, 1, 2)

	for _, v := range d {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func TestFib(t *testing.T) {
	n := Fib(6)
	if n != 8 {
		t.Error("测试失败")
	}
}
