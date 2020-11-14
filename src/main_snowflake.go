package main

import (
	"fmt"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
)

func main()  {
	node, err := utils.NewWorker(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(node.GetId())
}
