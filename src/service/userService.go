package service

import (
	"github.com/zhaiyjgithub/TagTalk-go/src/dao"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
)

type UserService interface {
	AddNewUser(user *model.User) error
	IsUserRegister(email string) bool
	GetUserByEmail(email string) *model.User
	GetNearByUsers(chatId string) []*model.User
	GetUserByChatID(chatId string) *model.User
	UpdateProfile(user *model.User) error
	UpdateImageWall(chatId string, name string) error
	GetImageWall(chatId string) []string
	UpdateTags(chatId string, names string) error
	GetTags(chatId string) *model.Tag
	UpdateAvatar(chatId string, avatar string) error
	UpdateBasicProfile(chatId string, user *model.User) error
}

func NewUserService() UserService {
	return &userService{dao: dao.NewUserDao(database.InstanceMysqlDB())}
}

type userService struct {
	dao *dao.UserDao
}

func (s *userService) AddNewUser(user *model.User) error {
	return s.dao.AddNewUser(user)
}

func (s *userService) IsUserRegister(email string) bool  {
	return s.dao.IsUserRegister(email)
}

func (s *userService) GetUserByEmail(email string) *model.User  {
	return s.dao.GetUserByEmail(email)
}

func (s *userService) GetNearByUsers(chatId string) []*model.User {
	return s.dao.GetNearByUsers(chatId)
}

func (s *userService) GetUserByChatID(chatId string) *model.User  {
	return s.dao.GetUserByChatID(chatId)
}

func (s *userService)UpdateProfile(user *model.User) error  {
	return s.dao.UpdateProfile(user)
}

func (s *userService) UpdateImageWall(chatId string, name string) error  {
	return s.dao.UpdateImageWall(chatId, name)
}

func (s *userService) GetImageWall(chatId string) []string  {
	return s.dao.GetImageWall(chatId)
}

func (s *userService) UpdateTags(chatId string, names string) error  {
	return s.dao.UpdateTags(chatId, names)
}

func (s *userService) GetTags(chatId string) *model.Tag  {
	return s.dao.GetTags(chatId)
}

func (s *userService) UpdateAvatar(chatId string, avatar string) error  {
	return s.dao.UpdateAvatar(chatId, avatar)
}

func (s *userService) UpdateBasicProfile(chatId string, user *model.User) error  {
	return s.dao.UpdateBasicProfile(chatId, user)
}