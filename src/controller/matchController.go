package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"github.com/zhaiyjgithub/TagTalk-go/src/service"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
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

func (c *MatchViewController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(iris.MethodPost, utils.GetMatchList,"GetMatchList")
	b.Handle(iris.MethodPost, utils.AddLikeOrDisLike, "AddLikeOrDisLike")
}

func (c *MatchViewController)GetMatchList()  {
	type Param struct {
		ChatId int64
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

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

func (c *MatchViewController) AddLikeOrDisLike()  {
	type Param struct {
		ChatId int64
		PeerChatId int64
		Type LikeType
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	key := ""
	if p.Type == Like {
		key = fmt.Sprintf("like_%d", p.ChatId)
	}else if p.Type == DisLike {
		key = fmt.Sprintf("dislike_%d", p.ChatId)
	}else if p.Type == Star {
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
	code, err := rd.SAdd(contextBg, key, item).Result()
	fmt.Printf("code: %d -- err: %v", code, err)
}

