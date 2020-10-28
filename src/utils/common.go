package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"time"
)

var defaultValidator = validator.New()
var Jwt, _ = jwt.New(15*time.Minute, jwt.HS256, []byte("hello"))

func ValidateParam(ctx iris.Context, param interface{}) error {
	err := ctx.ReadJSON(param)
	if  err != nil {
		response.Fail(ctx, response.Error, response.ParamErr, nil)
		return  err
	}

	err = defaultValidator.Struct(param)
	if err != nil {
		response.Fail(ctx, response.Error, err.Error(), nil)
		return err
	}

	return nil
}