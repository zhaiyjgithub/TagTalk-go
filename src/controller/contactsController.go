package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
)

type ContactsController struct {
	Ctx iris.Context
}

func (c *ContactsController) BeforeActivation(b mvc.BeforeActivation)  {

}

func (c *ContactsController) GetContacts()  {
	type Param struct {
		ChatID int64
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}



}
