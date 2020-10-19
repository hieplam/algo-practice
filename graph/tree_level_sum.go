package graph

import (
	"fmt"
	"strings"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func TreeLevelSum(t tree) string {
	var r []int
	levelNode := []tree{t}

	for len(levelNode) > 0 {
		sum := 0
		var childLevelNode []tree
		for _, v := range levelNode {
			sum += v.value
			if v.left != nil {
				childLevelNode = append(childLevelNode, *v.left)
			}
			if v.right != nil {
				childLevelNode = append(childLevelNode, *v.right)
			}
		}

		r = append(r, sum)
		levelNode = childLevelNode
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r)), ","), "[]")
}
