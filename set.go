package ACautomaton

var acMap *node

func SetList(list []string) {
	if len(list) == 0 {
		return
	}
	acMap = &node{}

	for _, words := range list {
		currNode := acMap
		for _, r := range words {
			if _, ok := (*currNode).child[r]; !ok {
				if (*currNode).child == nil {
					(*currNode).child = make(map[rune]*node)
				}
				(*currNode).child[r] = &node{}
			}
			currNode = (*currNode).child[r]
		}
		currNode.isAEnd = true
	}
}
