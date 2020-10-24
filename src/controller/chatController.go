package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ChatController struct {
	Ctx *iris.Context

}

func (c *ChatController) BeforeActivation(b mvc.BeforeActivation)  {
	maxSize := c.Ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := c.Ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		c.Ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	form := c.Ctx.Request().MultipartForm

	dir := "../web/sources/videos/"
	for _, hs := range form.File {
		if len(hs) > 0 {
			fh := hs[0]
			path := fmt.Sprintf("%s%s", dir, fh.Filename)
			_, _ = c.Ctx.SaveFormFile(fh, path)
		}
	}

	c.Ctx.WriteString("success")
}

func (c *ChatController) UploadFile()  {
	//f, h, err := c.c.Ctx.FormFile("file")
}