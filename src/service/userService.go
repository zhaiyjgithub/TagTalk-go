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
	GetNearByUsers(chatId int64) []*model.User
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

func (s *userService) GetNearByUsers(chatId int64) []*model.User {
	return s.dao.GetNearByUsers(chatId)
}