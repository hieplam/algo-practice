package general

import "fmt"

// RotateClockwiseMatrix /
/*

1 2 3		7 4 1
4 5 6  -->  8 5 2
7 8 9		9 6 3
    i  j            j  i
1 a[0][0]		1 a[0][2]
2 a[0][1]		2 a[1][2]
3 a[0][2]		3 a[2][2]

4 a[1][0]		4 a[0][1]
5 a[1][1]  		5 a[1][1]
6 a[1][2]		6 a[2][1]
*/
func RotateClockwiseMatrix(m [][]int) ([][]int, error) {
	r := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Println(r[i][j])
			r[j][len(m)-1-i] = m[i][j]
		}
	}
	return r, nil
}
