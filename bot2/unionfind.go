package main

type unionFind struct {
	parent map[cell]cell
	rank   map[cell]int
}

type cell struct {
	row int
	col int
}

func (c cell) edge(color color) int {
	if color == white {
		// White goes from left to right
		return c.col
	} else {
		// Black goes from top to bottom
		return c.row
	}
}

func newUnionFind() *unionFind {
	return &unionFind{parent: make(map[cell]cell), rank: make(map[cell]int)}
}

func (f *unionFind) find(x cell) cell {
	if px, ok := f.parent[x]; ok {
		if px == x {
			return x
		}
		gx := f.parent[px]
		if gx == px {
			return px
		}
		f.parent[x] = gx
		return f.find(gx)
	}

	f.parent[x] = x
	f.rank[x] = 0
	return x
}

func (f *unionFind) join(x, y cell) bool {
	repX, repY := f.find(x), f.find(y)
	if repX == repY {
		return false
	}

	if f.rank[repX] < f.rank[repY] {
		f.parent[repX] = repY
	} else if f.rank[repX] > f.rank[repY] {
		f.parent[repY] = repX
	} else {
		f.parent[repX] = repY
		f.rank[repY] += 1
	}

	return true
}

func (f *unionFind) connected(x, y cell) bool { return f.find(x) == f.find(y) }
