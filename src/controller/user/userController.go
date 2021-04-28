package user

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"github.com/zhaiyjgithub/TagTalk-go/src/response"
	"github.com/zhaiyjgithub/TagTalk-go/src/service"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	PinCacheTimeout = 60*time.Second
	AvatarBaseDir = "./src/web/source/avatar/"
	ImageWallsBaseDir = "./src/web/source/imagewalls/"
)

var contextBg = context.Background()

type Controller struct {
	Ctx iris.Context
	UserService service.UserService
}

func (c *Controller) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(iris.MethodPost, utils.RegisterNewDoctor,"RegisterNewDoctor")
	b.Handle(iris.MethodPost, utils.SendSignUpPin,"SendSignUpPin")
	b.Handle(iris.MethodPost, utils.Login, "Login")
	b.Handle(iris.MethodPost, utils.GetUserInfo, "GetUserInfo")
	b.Handle(iris.MethodPost, utils.UpdateProfile, "UpdateProfile")
	b.Handle(iris.MethodPost, utils.UploadImageWalls, "UploadImageWalls")
	b.Handle(iris.MethodGet, utils.ImageWalls, "ImageWalls")
	b.Handle(iris.MethodGet, utils.Avatar, "Avatar")
	b.Handle(iris.MethodPost, utils.GetImageWall, "GetImageWall")
	b.Handle(iris.MethodPost, utils.GetTags, "GetTags")
	b.Handle(iris.MethodPost, utils.UpdateTags, "UpdateTags")
	b.Handle(iris.MethodPost, utils.UpdateAvatar, "UpdateAvatar")
}

func (c *Controller) RegisterNewDoctor()  {
	type param struct {
		Name string `validate:"gt=0"`
		Email string `validate:"email"`
		Password string `validate:"min=6,max=20"`
		Pin string `validate:"len=6"`
	}

	var p param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	isExist := c.UserService.IsUserRegister(p.Email)
	if isExist {
		response.Fail(c.Ctx, response.IsExist, "Email has been registered", nil)
		return
	}

	pin , err := getSignUpPinFromCache(p.Email)
	if err != nil || pin != p.Pin {
		response.Fail(c.Ctx, response.Error, "Verification code is invalid", nil)
	}else {
		node, err := utils.NewWorker(1)
		if err != nil {
			response.Fail(c.Ctx, response.Error, err.Error(), nil)
			return
		}

		user := &model.User{
			Email: p.Email,
			Password: p.Password,
			Name: p.Name,
			ChatID: string(node.GetId()), //通过雪花算法生成唯一ID，作为Chat ID
		}

		err = c.UserService.AddNewUser(user)
		if err != nil {
			response.Fail(c.Ctx, response.Error, err.Error(), nil)
		}else {
			response.Success(c.Ctx, response.Successful, nil)
		}
	}
}

func (c *Controller) Login()  {
	type Param struct {
		Email string `validate:"email"`
		Password string `validate:"min=6,max=20"`
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	user := c.UserService.GetUserByEmail(p.Email)
	if user == nil {
		response.Fail(c.Ctx, response.Error, "", nil)
	}else {
		type UserInfo struct {
			*model.User
			Token string
		}

		token, _ := generateToken()

		var info UserInfo
		info.User = user
		info.Token = token
		response.Success(c.Ctx, response.Successful, &info)
	}
}

func (c *Controller) GetUserInfo()  {
	type Param struct {
		ChatID string
	}
	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}
	user:= c.UserService.GetUserByChatID(p.ChatID)
	response.Success(c.Ctx, response.Successful, &user)
}

func (c *Controller) SendSignUpPin()  {
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

	pin := generateSignUpPin()
	err = addSignUpPinToCache(p.Email, pin)
	
	err = utils.SendPinEmail(p.Email, pin)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}

	fmt.Printf("\r\n your pin: %s \r\n", pin)
}

func (c *Controller) UpdateProfile()  {
	maxSize := c.Ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := c.Ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		c.Ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	form := c.Ctx.Request().MultipartForm
	
	chatIdValues := form.Value["ChatID"]
	genderValues := form.Value["Gender"]
	bioValues := form.Value["Bio"]

	if len(chatIdValues[0]) == 0 || len(genderValues[0]) == 0 || len(bioValues[0]) == 0 {
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
		return
	}

	chatId := chatIdValues[0]
	gender := genderValues[0]
	bio := bioValues[0]

	if chatIdValues == nil || genderValues == nil || bioValues == nil {
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
		return
	}

	if len(bio) < 20 || len(bio) > 150 {
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
	}

	fullName := ""
	for _, hs := range form.File {
		if len(hs) > 0 {
			fh := hs[0]

			id, _ := strconv.Atoi(chatId)
			encodeFileName := utils.GenerateFileName(id)
			ext := filepath.Ext(fh.Filename)
			fullName = fmt.Sprintf("%s%s", encodeFileName, ext)
			path := fmt.Sprintf("%s%s", AvatarBaseDir, fullName)
			_, err = c.Ctx.SaveFormFile(fh, path)
		}
	}

	u := &model.User{}
	u.ChatID = chatId

	u.Gender = model.Male
	if gender != "1" {
		u.Gender = model.Female
	}
	u.Bio = bio
	u.Avatar = fullName

	err = c.UserService.UpdateProfile(u)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	}

	if err != nil {
		response.Fail(c.Ctx, response.Error, "Upload image failed", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) Avatar()  {
	name := c.Ctx.URLParam("name")
	path := fmt.Sprintf(AvatarBaseDir + name)
	_ = c.Ctx.ServeFile(path)
}

func (c *Controller) ImageWalls()  {
	name := c.Ctx.URLParam("name")
	path := fmt.Sprintf(ImageWallsBaseDir + name)
	_ = c.Ctx.ServeFile(path)
}

func (c *Controller) GetImageWall()  {
	type Param struct {
		ChatID string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	names := c.UserService.GetImageWall(p.ChatID)
	response.Success(c.Ctx, response.Successful, names)
}

func (c *Controller) UploadImageWalls()  {
	maxSize := c.Ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := c.Ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		c.Ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	form := c.Ctx.Request().MultipartForm

	chatIdValues := form.Value["ChatID"]
	newImagesValues := form.Value["NewImageNames"]
	deleteImagesValues := form.Value["DeleteImageNames"]

	if chatIdValues == nil || newImagesValues == nil || deleteImagesValues == nil {
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
		return
	}

	if len(chatIdValues[0]) == 0 || len(newImagesValues) == 0 || len(deleteImagesValues) == 0 {
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
		return
	}

	if len(newImagesValues) == 0 {
		response.Fail(c.Ctx, response.Error, response.ParamErr, nil)
		return
	}

	chatId := chatIdValues[0]
	err = c.UserService.UpdateImageWall(chatId, newImagesValues[0])

	deleteImages := strings.Split(deleteImagesValues[0], ",")
	fmt.Println(deleteImages)

	for _, fileName := range deleteImagesValues {
		if len(fileName) == 0 {
			continue
		}

		path := fmt.Sprintf("%s%s", ImageWallsBaseDir, fileName)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("file does not exist")
			continue
		}

		err := os.Remove(path)
		if err != nil {
			fmt.Printf("remove %s failed \r\n")
		}
	}

	for _, hs := range form.File {
		if len(hs) > 0 {
			fh := hs[0]
			path := fmt.Sprintf("%s%s", ImageWallsBaseDir, fh.Filename)
			_, err = c.Ctx.SaveFormFile(fh, path)
		}
	}
	if err != nil {
		response.Fail(c.Ctx, response.Error, "Upload image failed", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
 }

func (c *Controller) UpdateTags()  {
	type Param struct {
		ChatID string
		Names string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	err = c.UserService.UpdateTags(p.ChatID, p.Names)
	if err != nil {
		response.Fail(c.Ctx, response.Error, "", nil)
	}else {
		response.Success(c.Ctx, response.Successful, nil)
	}
}

func (c *Controller) GetTags()  {
	type Param struct {
		ChatID string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	ts := c.UserService.GetTags(p.ChatID)
	response.Success(c.Ctx, response.Successful, ts)
}

func generateToken() (string, error) {
	var claims jwt.Claims
	token, err := utils.Jwt.Token(claims)

	if err != nil {
		return "", err
	}

	return token, nil
}

func generateSignUpPin() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return pin
}

func addSignUpPinToCache(key string, val string) error {
	rdb := database.InstanceRedisDB()
	return rdb.Set(contextBg, key, val, PinCacheTimeout).Err()
}

func getSignUpPinFromCache(key string) (string, error) {
	rdb := database.InstanceRedisDB()
	val, err := rdb.Get(contextBg, key).Result()
	return val, err
}
