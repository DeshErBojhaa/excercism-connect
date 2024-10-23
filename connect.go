// Package connect
// DSU solution i.e. union find solution.
package connect

type Data struct {
	ids                      map[int]map[int]int
	top, bottom, left, right int
}

func (d Data) unionFind(lines []string, piece uint8, source, dest int) bool {
	N := d.right
	parent := make([]int, N+1) // d.right has the max value
	rank := make([]int, N+1)
	for i := range parent {
		parent[i] = i
		rank[i] = 0
	}
	for i, s := range lines {
		for j, c := range s {
			if byte(c) != piece {
				continue
			}
			id := d.ids[i][j]
			// Special case top
			if i == 0 && piece == 'O' {
				join(id, d.top, parent, rank)
			}
			// Special case bottom
			if i == len(lines)-1 && piece == 'O' {
				join(id, d.bottom, parent, rank)
			}
			// Special case left
			if j == 0 && piece == 'X' {
				join(id, d.left, parent, rank)
			}
			// Special case right
			if j == len(s)-1 && piece == 'X' {
				join(id, d.right, parent, rank)
			}

			for _, dx := range []int{0, 1} {
				for _, dy := range []int{-1, 0, 1} {
					if (dx == 0 && dy == 0) || (dx == 1 && dy == 1) {
						continue
					}
					ni, nj := i+dx, j+dy
					if ni < 0 || ni >= len(lines) || nj < 0 || nj >= len(s) {
						continue
					}
					if lines[ni][nj] != piece {
						continue
					}
					nid := d.ids[ni][nj]
					join(id, nid, parent, rank)
				}

			}
		}
	}

	return findParent(source, parent) == findParent(dest, parent)
}

func join(a, b int, parent, rank []int) {
	pa, pb := findParent(a, parent), findParent(b, parent)
	if pa == pb {
		return
	}
	if rank[pa] < rank[pb] {
		pa, pb = pb, pa
	}
	parent[pb] = pa
	if rank[pa] == rank[pb] {
		rank[pa]++
	}
}

func findParent(now int, parent []int) int {
	if parent[now] == now {
		return now
	}
	parent[now] = findParent(parent[now], parent)
	return parent[now]
}

func ResultOf(lines []string) (string, error) {
	n, m, id := len(lines), len(lines[0]), 1
	ids := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := ids[i]; !ok {
			ids[i] = make(map[int]int)
		}
		for j := 0; j < m; j++ {
			ids[i][j] = id
			id++
		}
	}

	top, bottom, left, right := id+1, id+2, id+3, id+4
	d := Data{
		ids: ids,
		top: top, bottom: bottom, left: left, right: right,
	}
	if d.unionFind(lines, 'O', top, bottom) {
		return "O", nil
	}
	if d.unionFind(lines, 'X', left, right) {
		return "X", nil
	}
	return "", nil
}
