package ACautomaton

type node struct {
	isAEnd bool           //是否是一个词语的结束点
	child  map[rune]*node //子节点指针
}
