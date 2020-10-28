package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
)

var defaultValidator = validator.New()

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