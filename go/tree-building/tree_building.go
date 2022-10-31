package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Records []Record

func (r Records) Less(a, b int) bool {
	return r[a].ID < r[b].ID
}

func (r Records) Swap(a, b int) {
	r[a], r[b] = r[b], r[a]
}

func (r Records) Len() int {
	return len(r)
}

type Node struct {
	ID       int
	Children []*Node
	//Parent   *Node
}

type Sibilings []*Node

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(Records(records))

	if records[len(records)-1].ID != len(records)-1 {
		return nil, errors.New("non-continuous")
	}

	var root *Node
	siblings := make(map[int]Sibilings)
	rootFound := false

	// Group by parent.
	for _, r := range records {
		if r.ID == 0 && r.Parent == 0 {
			if rootFound {
				return nil, errors.New("duplicate root")
			}
			rootFound = true
			root = &Node{r.ID, nil}
			continue
		}

		if r.ID == r.Parent {
			return nil, errors.New("cycle directly")
		}

		if r.ID < r.Parent {
			return nil, errors.New("cycle indirectly")
		}

		siblings[r.Parent] = append(siblings[r.Parent], &Node{r.ID, nil})
	}

	if !rootFound {
		return nil, errors.New("root not found")
	}

	_, err := appendChildren(root, siblings)
	if err != nil {
		return nil, err
	}

	return root, nil
}

func appendChildren(node *Node, siblings map[int]Sibilings) (*Node, error) {
	if siblings[node.ID] == nil {
		return node, nil
	}

	children := siblings[node.ID]

	node.Children = children

	for i, child := range children {
		dup := sort.Search(len(children), func(j int) bool {
			// same child.
			if i == j {
				return false
			}
			// duplicate child.
			return children[j].ID == child.ID
		})

		if dup != len(children) {
			return nil, errors.New("duplicate record")
		}

		appendChildren(child, siblings)
	}

	return node, nil
}
