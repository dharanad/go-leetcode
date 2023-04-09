package main

import "fmt"

var (
	X = []int{0, 1, -1, 0}
	Y = []int{1, 0, 0, -1}
)

func markVisisted(r, c, rows, cols int, grid [][]int) {
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 1 {
		return
	}
	grid[r][c] = 1
	for i := 0; i < 4; i++ {
		nr := r + X[i]
		nc := c + Y[i]
		markVisisted(nr, nc, rows, cols, grid)
	}
}

func closedIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	for j := 0; j < cols; j++ {
		if grid[0][j] == 0 {
			markVisisted(0, j, rows, cols, grid)
		}
		if grid[rows-1][j] == 0 {
			markVisisted(rows-1, j, rows, cols, grid)
		}
	}
	for j := 0; j < rows; j++ {
		if grid[j][0] == 0 {
			markVisisted(j, 0, rows, cols, grid)
		}
		if grid[j][cols-1] == 0 {
			markVisisted(j, cols-1, rows, cols, grid)
		}
	}
	count := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 0 {
				count++
				markVisisted(i, j, rows, cols, grid)
			}
		}
	}
	return count
}

func findRedundantConnection(edges [][]int) []int {
	var res []int
	dsu := NewDSU(len(edges))
	for _, e := range edges {
		if dsu.IsUnion(e[0], e[1]) {
			res = e
		} else {
			dsu.Union(e[0], e[1])
		}
	}
	fmt.Println(dsu.parent)
	return res
}

type DSU struct {
	parent []int //parent[i] indicate parent of i the node
	// invariant u, v
	// u < v
	// u will the parent of v
}

func NewDSU(n int) *DSU {
	dsu := &DSU{}
	dsu.parent = make([]int, n+1)
	for i := 0; i <= n; i++ {
		dsu.parent[i] = i
	}
	return dsu
}

func (d *DSU) IsUnion(u, v int) bool {
	pu := d.Find(u)
	pv := d.Find(u)
	return pu == pv
}

func (d *DSU) Union(u, v int) {
	if v > u {
		u, v = v, u
	}
	d.parent[v] = d.parent[u]
}

func (d *DSU) Find(u int) int {
	if d.parent[u] == u { // parent[4] = 4
		return u
	}
	// parent[4] = 3 -> find(3)
	d.parent[u] = d.Find(d.parent[u])
	return d.parent[u]
}

type BoundaryChecker func(int, int) bool

func dfs(r, c int, grid [][]int, checker BoundaryChecker) int {
	if !checker(r, c) || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	count := 1
	for i := 0; i < 4; i++ {
		nr := r + X[i]
		nc := c + Y[i]
		count += dfs(nr, nc, grid, checker)
	}
	return count
}

func numEnclaves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var isSafe BoundaryChecker = func(r int, c int) bool {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}
		return true
	}
	for j := 0; j < cols; j++ {
		if grid[0][j] == 1 {
			dfs(0, j, grid, isSafe)
		}
		if grid[rows-1][j] == 1 {
			dfs(rows-1, j, grid, isSafe)
		}
	}
	for j := 0; j < rows; j++ {
		if grid[j][0] == 1 {
			dfs(j, 0, grid, isSafe)
		}
		if grid[j][cols-1] == 1 {
			dfs(j, cols-1, grid, isSafe)
		}
	}
	res := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 1 {
				res += dfs(i, j, grid, isSafe)
			}
		}
	}
	return res
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return cloneGraphHelper(node, visited)
}

func cloneGraphHelper(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return node
	}
	if _, ok := visited[node]; ok {
		return visited[node]
	}

	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, cloneGraphHelper(n, visited))
	}
	return newNode
}
