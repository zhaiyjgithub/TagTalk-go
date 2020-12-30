package dao

import (
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"gorm.io/gorm"
)

type ContactsDao struct {
	engine *gorm.DB
}

func NewContactsDao(engine *gorm.DB) *ContactsDao {
	return &ContactsDao{engine: engine}
}

func (d *ContactsDao) GetContactsByChatID(chatId int64) []*model.Contact {
	var list []*model.Contact
	d.engine.Where("chat_id = ?", chatId).Find(&list)

	return list
}

func (d *ContactsDao) AddNewFriend(contact *model.Contact) error {
	db := d.engine.Create(contact)
	return db.Error
}