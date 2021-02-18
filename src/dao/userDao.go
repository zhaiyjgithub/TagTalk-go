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

func (d *UserDao) GetNearByUsers(chatId string) []*model.User {
	var list []*model.User
	_ = d.engine.Where("chat_id != ?", chatId).Limit(10).Find(&list)

	return list
}

func (d *UserDao) GetUserByChatID(chatId string) *model.User  {
	var u *model.User
	_ = d.engine.Where("chat_id = ?", chatId).Limit(1).Find(&u)
	return u
}

func (d *UserDao) UpdateProfile(user *model.User) error  {
	db := d.engine.Model(&model.User{}).Where("chat_id = ?", user.ChatID).Updates(model.User{Gender: user.Gender, Bio: user.Bio})

	return db.Error
}