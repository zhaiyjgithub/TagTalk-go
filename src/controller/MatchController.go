package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"github.com/zhaiyjgithub/TagTalk-go/src/service"
)

type MatchViewController struct {
	Ctx iris.Context
	UserService service.UserService
}

type LikeType int
const (
	Like LikeType = 0
	DisLike LikeType = 1
	Star LikeType = 2
)

func (c *MatchViewController)GetMatchList()  {
	type Param struct {
		ChatId int64
	}

	var p Param

	type MatchViewModel struct {
		User *model.User
		Likes []string
		DisLikes []string
		Stars []string
	}

	users := c.UserService.GetNearByUsers(p.ChatId)

	likesKey:= fmt.Sprintf("like_%d", p.ChatId)
	disLikesKey := fmt.Sprintf("dislike_%d", p.ChatId)
	starLikesKey := fmt.Sprintf("starLike_%d", p.ChatId)
	var vModels []*MatchViewModel
	for _, user := range users{
		likes := getLikesOrDisLikes(likesKey)
		disLikes := getLikesOrDisLikes(disLikesKey)
		starLikes := getLikesOrDisLikes(starLikesKey)

		vm := &MatchViewModel{
			User: user,
			Likes: likes,
			DisLikes: disLikes,
			Stars: starLikes,
		}

		vModels = append(vModels, vm)
	}

	response.Success(c.Ctx, response.Successful, &vModels)
}

func (c *MatchViewController) AddLike()  {
	type Param struct {
		ChatId int64
		PeerChatId int64
		likeType LikeType
	}

	var p Param

	key := ""
	if p.likeType == Like {
		key = fmt.Sprintf("like_%d", p.ChatId)
	}else if p.likeType == DisLike {
		key = fmt.Sprintf("dislike_%d", p.ChatId)
	}else if p.likeType == Star {
		likeKey := fmt.Sprintf("like_%d", p.ChatId)
		addLikesOrDisLikes(likeKey, fmt.Sprintf("%d", p.PeerChatId))

		key = fmt.Sprintf("starLike_%d", p.ChatId)
	}
	addLikesOrDisLikes(key, fmt.Sprintf("%d", p.PeerChatId))

	response.Success(c.Ctx, response.Successful, nil)
}

func getLikesOrDisLikes(key string) []string{
	rd := database.InstanceRedisDB()
	items, _ := rd.SMembers(contextBg, key).Result()
	return items
}

func addLikesOrDisLikes(key string, item string) {
	rd := database.InstanceRedisDB()
	rd.SAdd(contextBg, key, item)
}

