package trees

type tree struct {
	value       int
	left, right *tree
}

func nodeDepths(r *tree) int {
	d := 1
	s := 0
	arr := []*tree{r}
	childArr := []*tree{}

	for len(arr) > 0 {
		n := arr[0]
		arr = arr[1:]

		if n.left != nil {
			s += d
			childArr = append(childArr, n.left)
		}
		if n.right != nil {
			s += d
			childArr = append(childArr, n.right)
		}

		if len(arr) == 0 {
			d++
			arr = childArr
			childArr = []*tree{}
		}
	}

	return s
}
