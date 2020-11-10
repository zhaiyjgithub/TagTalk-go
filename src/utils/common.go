package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/zhaiyjgithub/TagTalk-go/src/conf"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"gopkg.in/gomail.v2"
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

func SendPinEmail(email string, pin string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.ServerEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "DrFinder Verification code")

	body := fmt.Sprintf("Your Tag Talk pin: %s", pin)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(conf.Smtp, 587, conf.ServerEmail, conf.ServerEmailPwd)

	return d.DialAndSend(m)
}