package main

import "fmt"

type unionFind struct {
	root []int
}

func (u *unionFind) Find(val int) int {
	return u.root[val]
}

func (u *unionFind) Union(x, y int) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return
	}
	for i, val := range u.root {
		if rootY != val {
			continue
		}
		u.root[i] = rootX
	}
}

func (u *unionFind) Connected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func NewUnionFind(size int) *unionFind {
	uf := &unionFind{
		root: make([]int, size),
	}

	for i := range uf.root {
		uf.root[i] = i
	}

	return uf
}

func main() {
	uf := NewUnionFind(10)
	// 1-2-5-6-7 3-8-9 4
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)

	fmt.Println(uf.Connected(1, 5)) // true
	fmt.Println(uf.Connected(5, 7)) // true
	fmt.Println(uf.Connected(4, 9)) // false

	// 1-2-5-6-7 3-8-9-4
	uf.Union(9, 4)
	fmt.Println(uf.Connected(4, 9)) // true
}
