package main

import (
	"fmt"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
)

func main()  {
	s := utils.NewSet()
	s.Add("hello")
	//
	//if s.Contains("hello") {
	//	fmt.Println("contain")
	//}else {
	//	fmt.Println("not contain")
	//}

	fmt.Println(s.Values())
	s.Delete("hello")
	fmt.Println(s.Values())
}
