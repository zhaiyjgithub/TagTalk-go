package service

import (
	"github.com/zhaiyjgithub/TagTalk-go/src/dao"
	"github.com/zhaiyjgithub/TagTalk-go/src/database"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
)

type ContactsService interface {
	GetContactsByChatID(chatId int64) []*model.Contact
	AddNewFriend(contact *model.Contact) error
}

type contactsService struct {
	dao *dao.ContactsDao
}

func NewContactsService() ContactsService  {
	return &contactsService{dao: dao.NewContactsDao(database.InstanceMysqlDB())}
}

func (s *contactsService) GetContactsByChatID(chatId int64) []*model.Contact {
	return s.dao.GetContactsByChatID(chatId)
}

func (s *contactsService) AddNewFriend(contact *model.Contact) error  {
	return s.AddNewFriend(contact)
}
