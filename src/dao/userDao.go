package dao

import (
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"gorm.io/gorm"
)

type UserDao struct {
	engine *gorm.DB
}

func NewUserDao(engine *gorm.DB) *UserDao {
	return &UserDao{engine: engine}
}

func (d *UserDao) AddNewUser(user *model.User) error {
	db := d.engine.Create(user)
	return db.Error
}

func (d *UserDao) IsUserRegister(email string) bool {
	var u model.User
	db := d.engine.Where("email = ?", email).Find(&u)

	if db.Error != nil {
		return true
	}

	return u.ID > 0
}
