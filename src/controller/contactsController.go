package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"github.com/zhaiyjgithub/TagTalk-go/src/service"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
	"net/http"
)

type ContactsController struct {
	Ctx iris.Context
	ContactsService service.ContactsService
}

func (c *ContactsController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(http.MethodPost, utils.GetContactsByChatID, "GetContactsByChatID")
	b.Handle(http.MethodPost, utils.AddNewFriend, "AddNewFriend")
}

func (c *ContactsController) AddNewFriend()  {
	type Param struct {
		ChatID int64
		FriendID int64
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	m := &model.Contact{ChatID: p.ChatID, FriendID: p.FriendID}
	err = c.ContactsService.AddNewFriend(m)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *ContactsController) GetContactsByChatID()  {
	type Param struct {
		ChatID int64
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	list := c.ContactsService.GetContactsByChatID(p.ChatID)
	response.Success(c.Ctx, response.Successful, list)
}
