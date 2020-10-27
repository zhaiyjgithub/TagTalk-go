package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
)

type UserController struct {
	Ctx *iris.Context
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(iris.MethodPost, utils.RegisterNewDoctor,"RegisterNewDoctor")
}

func (c *UserController) RegisterNewDoctor()  {
	type param struct {
		Name string `validate:"gt=0"`
		Email string `validate:"email"`
		Password string `validate:"min=6,max=20"`
		Gender model.GenderType
		Code string `validate:"len=6"`
	}
}
