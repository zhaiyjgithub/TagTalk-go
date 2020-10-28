package controller

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"github.com/zhaiyjgithub/TagTalk-go/src/service"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
	"math/rand"
	"time"
)

const (
	VerificationCodeTimeout time.Duration = 0
)

var contextBg = context.Background()

type UserController struct {
	Ctx iris.Context
	UserService service.UserService
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(iris.MethodPost, utils.RegisterNewDoctor,"RegisterNewDoctor")
	b.Handle(iris.MethodPost, utils.RequestVerificationCode,"RequestVerificationCode")
}

func (c *UserController) RegisterNewDoctor()  {
	type param struct {
		Name string `validate:"gt=0"`
		Email string `validate:"email"`
		Password string `validate:"min=6,max=20"`
		Gender model.GenderType
		Code string `validate:"len=6"`
	}

	var p param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	isExist := c.UserService.IsUserRegister(p.Email)
	if isExist {
		response.Fail(c.Ctx, response.IsExist, "Email is registered", nil)
		return
	}

	code , err := getVerificationCodeToRedis(p.Email)
	if err != nil || code != p.Code {
		response.Fail(c.Ctx, response.Error, "Verification code is invalid", nil)
	}else {
		user := &model.User{
			Email: p.Email,
			Password: p.Password,
			Name: p.Name,
			Gender: p.Gender,
		}

		err = c.UserService.AddNewUser(user)
		if err != nil {
			response.Fail(c.Ctx, response.Error, err.Error(), nil)
		}else {
			response.Success(c.Ctx, response.Successful, nil)
		}
	}
}

func (c *UserController) RequestVerificationCode()  {
	type Param struct {
		Email string `validate:"email"`
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	isExist := c.UserService.IsUserRegister(p.Email)
	if isExist {
		response.Fail(c.Ctx, response.IsExist, "Email is registered", nil)
		return
	}

	code := generateVerificationCode()
	err = addVerificationCodeToRedis(p.Email, code)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}

	fmt.Printf("your code: %s \r\n", code)
}

func generateVerificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}

func addVerificationCodeToRedis(key string, val string) error {
	rdb := database.InstanceRedisDB()
	return rdb.Set(contextBg, key, val, VerificationCodeTimeout).Err()
}

func getVerificationCodeToRedis(key string) (string, error) {
	rdb := database.InstanceRedisDB()
	val, err := rdb.Get(contextBg, key).Result()
	return val, err
}

