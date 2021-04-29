package user

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
	"path/filepath"
	"strconv"
)

func (c *Controller) UpdateAvatar()  {
	maxSize := c.Ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := c.Ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		c.Ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	form := c.Ctx.Request().MultipartForm
	chatIdValues := form.Value["ChatID"]
	if chatIdValues == nil || len(chatIdValues[0]) == 0{
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
		return
	}

	chatId := chatIdValues[0]
	for _, hs := range form.File {
		if len(hs) > 0 {
			fh := hs[0]
			id, _ := strconv.Atoi(chatId)
			encodeFileName := utils.GenerateFileName(id)
			ext := filepath.Ext(fh.Filename)
			fullName:= fmt.Sprintf("%s%s", encodeFileName, ext)
			path := fmt.Sprintf("%s%s", AvatarBaseDir, fullName)
			_, err = c.Ctx.SaveFormFile(fh, path)
			if err == nil {
				err = c.UserService.UpdateAvatar(chatId, fullName)
			}
		}
	}
	if err != nil {
		response.Fail(c.Ctx, response.Error, "Upload image failed", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) UpdateBasicProfile()  {
	type Param struct {
		ChatID string
		Name string
		Bio string
		Gender model.GenderType
	}

	var p Param
	err:= utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	user:= &model.User{Name: p.Name, Bio: p.Bio, Gender: p.Gender}
	err = c.UserService.UpdateBasicProfile(p.ChatID, user)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "Upload image failed", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}