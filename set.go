package ACautomaton

var acMap *node

//输入需要匹配的内容列表
func SetList(list []string) {
	acMap = &node{}
	if len(list) == 0 {
		return
	}

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
