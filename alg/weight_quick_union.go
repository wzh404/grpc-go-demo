package alg

// 如果一棵大树和一棵小树合并，避免将大树放在小树的下面，就可以一定程度上避免更高的树

var (
	w = []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

func WQUnion(p, q int32) {
	m := root(p)
	n := root(q)

	if w[m] >= w[n] {
		a[n] = m
		w[m] += w[n]
	} else {
		a[m] = n
		w[n] += w[m]
	}
}
