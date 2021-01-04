package model

type Contact struct {
	ID     int   `gorm:"column:id;primary_key" json:"-"`
	ChatID string `gorm:"column:chat_id"`
	FriendID string `gorm:"column:friend_id"`
}

// TableName sets the insert table name for this struct type
func (c *Contact) TableName() string {
	return "contacts"
}
