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

func (d *UserDao) GetUserByEmail(email string) *model.User {
	var u model.User
	db := d.engine.Where("email = ? ", email).Find(&u)
	if db.Error != nil  || u.ID == 0{
		return nil
	}

	return &u
}

func (d *UserDao) GetNearByUsers(chatId int64) []*model.User {
	var list []*model.User
	_ = d.engine.Where("email != ?", chatId).Limit(10).Find(&list)

	return list
}