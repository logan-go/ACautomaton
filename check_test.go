package ACautomaton

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	list := []string{
		"习大大",
		"习惯",
		"低端人口",
	}
	SetList(list)

	msgList := []string{
		"政府说需要把低端人口驱逐出北京，是真的么？",
		"政府说需要把低端人驱逐出北京，是真的么？",
		"习大大可没有说过这样的话",
		"总理说习大了么？",
	}

	for _, msg := range msgList {
		fmt.Println(msg, Check(msg))
	}
}
