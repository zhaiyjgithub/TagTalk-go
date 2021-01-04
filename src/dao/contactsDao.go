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

func (d *ContactsDao) GetContactsByChatID(chatId string) []*model.User {
	var contacts []*model.Contact
	d.engine.Raw("SELECT chat_id, friend_id FROM contacts WHERE chat_id = ? UNION " +
		"SELECT chat_id, friend_id FROM contacts WHERE friend_id = ?", chatId, chatId).Scan(&contacts)

	var friendIds []string
	for _, contact := range contacts {
		friendIds = append(friendIds, contact.FriendID)
	}

	var friends []*model.User
	d.engine.Raw("SELECT * FROM users WHERE chat_id IN (?)", friendIds).Scan(&friends)

	return friends
}

func (d *ContactsDao) AddNewFriend(contact *model.Contact) error {
	db := d.engine.Create(contact)
	return db.Error
}