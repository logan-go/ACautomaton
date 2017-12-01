package ACautomaton

import (
	"fmt"
	"strings"
	"testing"
)

func TestSetList(t *testing.T) {
	list := []string{
		"习大大",
		"习大大",
		"习大大FDASFDSAF大大大多",
		"习惯",
		"FUCK",
		"FK",
		"FU",
		"F**k",
	}
	SetList(list)
	Printr(acMap, 0)
}

func Printr(currNode *node, deep int) {
	if currNode.child != nil {
		for k, v := range currNode.child {
			fmt.Println(strings.Repeat("\t", deep), fmt.Sprintf("%c", rune(k)), v.isAEnd)
			if v.child != nil {
				Printr(v, deep+1)
			}
		}
	}
}
