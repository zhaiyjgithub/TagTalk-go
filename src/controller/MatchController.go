package controller

import "github.com/zhaiyjgithub/TagTalk-go/src/model"

func GetMatchList()  {
	type MatchViewModel struct {
		User *model.User
		Likes
	}
}
