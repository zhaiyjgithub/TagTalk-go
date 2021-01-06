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
	b.Handle(iris.MethodPost, utils.GetNearByPeople,"GetNearByPeople")
	b.Handle(iris.MethodPost, utils.AddLikeStatus, "AddLikeStatus")
}

func (c *MatchViewController)GetNearByPeople()  {
	type Param struct {
		ChatId string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	type MatchViewModel struct {//每个附近的人基本信息
		User *model.User
		Likes []string
		DisLikes []string
		Stars []string
	}

	//TO-DO
	//这里后面将使用Redis计算附近的人
	users := c.UserService.GetNearByUsers(p.ChatId)

	var vModels []*MatchViewModel
	for _, user := range users{
		likesKey:= getRedisKey(Like, user.ChatID)
		disLikesKey := getRedisKey(DisLike, user.ChatID)
		starLikesKey := getRedisKey(Star, user.ChatID)

		likes := getChatIDListFromRedisByKey(likesKey)
		disLikes := getChatIDListFromRedisByKey(disLikesKey)
		starLikes := getChatIDListFromRedisByKey(starLikesKey)

		vModel := &MatchViewModel{
			User: user,
			Likes: likes,
			DisLikes: disLikes,
			Stars: starLikes,
		}

		vModels = append(vModels, vModel)
	}

	response.Success(c.Ctx, response.Successful, &vModels)
}

func (c *MatchViewController) AddLikeStatus()  {
	type Param struct {
		ChatId string
		PeerChatId string
		Type LikeType
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	key := getRedisKey(p.Type, p.ChatId)
	 if p.Type == Star {
		likeKey := getRedisKey(Like, p.ChatId)
		addChatIDToRedis(likeKey, p.PeerChatId)
	}
	addChatIDToRedis(key, p.PeerChatId)

	response.Success(c.Ctx, response.Successful, nil)
}

func getRedisKey(likeType LikeType, chatId string) string {
	key := ""
	switch likeType {
	case Like:
		key = fmt.Sprintf("like_%s", chatId)
		break
	case DisLike:
		key = fmt.Sprintf("disLike_%s", chatId)
		break
	case Star:
		key = fmt.Sprintf("starLike_%s", chatId)
		break
	default:
	}
	
	return key
}

func getChatIDListFromRedisByKey(key string) []string{
	rd := database.InstanceRedisDB()
	items, _ := rd.SMembers(contextBg, key).Result()
	return items
}

func addChatIDToRedis(key string, item string) {
	rd := database.InstanceRedisDB()
	code, err := rd.SAdd(contextBg, key, item).Result()
	fmt.Printf("code: %d -- err: %v", code, err)
}

