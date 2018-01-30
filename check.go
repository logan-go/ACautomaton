package ACautomaton

func Check(msg string) bool {
	if acMap == nil {
		return false
	}
	m := []rune(msg)
	for index, word := range m {
		if _, ok := acMap.child[word]; ok {
			if check(m[index:]) {
				return true
			}
		}
	}
	return false
}

func check(msg []rune) bool {
	currNode := acMap
	for _, r := range msg {
		if v, ok := (*currNode).child[r]; ok {
			if v.isAEnd {
				return true
			}
		} else {
			return false
		}
		currNode = (*currNode).child[r]
	}
	return false
}
