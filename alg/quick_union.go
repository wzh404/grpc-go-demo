package alg

// Quick-Union 算法的缺点在于树可能太高了,需要回溯一棵瘦长的树

func root(i int32) int32 {
	for i != a[i] {
		i = a[i]
	}

	return i
}

func QUnion(p, q int32) {
	m := root(p)
	n := root(q)

	a[m] = n
}

func QUFind(p, q int32) bool {
	return root(p) == root(q)
}
