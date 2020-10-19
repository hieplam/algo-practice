package graph

//link: https://www.hackerrank.com/challenges/connected-cell-in-a-grid/problem
func ConnectedCell(m [][]int) int {
	max := 0
	for r, _ := range m {
		for c, _ := range m[r] {
			if m[r][c] == 0 {
				continue
			}

			region := getMaxRegion(m, r, c)
			if region > max {
				max = region
			}
		}

	}

	return max
}

func getMaxRegion(m [][]int, r int, c int) int {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[r]) {
		return 0
	}
	if m[r][c] == 0 {
		return 0
	}
	size := 1
	m[r][c] = 0 //marked
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i == r && j == c {
				continue
			}
			size += getMaxRegion(m, i, j)
		}
	}
	return size
}
