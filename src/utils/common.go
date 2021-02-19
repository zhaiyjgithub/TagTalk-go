package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/zhaiyjgithub/TagTalk-go/src/conf"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"gopkg.in/gomail.v2"
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


func GenerateFileName(chatId int) string {
	data := []byte(fmt.Sprintf("%d-%d", chatId, time.Now().Unix()))
	md5er := md5.New()
	md5er.Write(data)

	return hex.EncodeToString(md5er.Sum(nil))
}

