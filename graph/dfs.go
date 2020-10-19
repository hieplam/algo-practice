package graph

import "strings"

type node struct {
	name   string
	childs []node
}

func dfsRecursive(n node, r *[]string) {
	*r = append(*r, n.name)
	for _, v := range n.childs {
		dfsRecursive(v, r)
	}
}

func DFS(n node) string {
	var r []string
	dfsRecursive(n, &r)
	return strings.Join(r, ",")
}
