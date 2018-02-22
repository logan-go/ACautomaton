package ACautomaton

import "fmt"

func ExampleSetList() {
	list := []string{
		"你好",
		"Hello！",
		"自动机",
	}
	SetList(list)
}

func ExampleCheck() {
	str := "你好世界！"
	fmt.Println("匹配结果：", Check(str))
}
