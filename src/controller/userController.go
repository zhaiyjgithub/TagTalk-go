package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
)

type userController struct {
	Ctx *iris.Context
}



func (c *userController) registerNewUser()  {
	type param struct {
		Name string `validate:"gt=0"`
		Email string `validate:"email"`
		Password string `validate:"min=6,max=20"`
		Gender model.GenderType
		Code string `validate:"len=6"`
	}
}
