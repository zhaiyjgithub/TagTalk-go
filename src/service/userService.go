package service

import (
	"github.com/zhaiyjgithub/TagTalk-go/src/dao"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
)

type UserService interface {
	AddNewUser(user *model.User) error
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
