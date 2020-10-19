package graph

import "testing"

func TestDFS(t *testing.T) {
	testCases := []struct {
		name     string
		input    node
		expected string
	}{
		{
			name: "testcase 1",
			input: node{
				name: "A",
				childs: []node{
					node{name: "B", childs: []node{
						node{name: "E"},
						node{name: "F", childs: []node{
							node{name: "I"},
							node{name: "J"},
						}},
					}},
					node{name: "C"},
					node{name: "D", childs: []node{
						node{name: "G", childs: []node{node{name: "K"}}},
						node{name: "H"},
					}},
				},
			},
			expected: "A,B,E,F,I,J,C,D,G,K,H",
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			if r := DFS(v.input); r != v.expected {
				t.Errorf("actual: %v, expected: %v", r, v.expected)
			}
		})
	}
}
