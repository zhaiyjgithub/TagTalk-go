package model

type Contact struct {
	ID     int   `gorm:"column:id;primary_key"`
	ChatID int64 `gorm:"column:chat_id"`
	PeerID int64 `gorm:"column:peer_id"`
}

// TableName sets the insert table name for this struct type
func (c *Contact) TableName() string {
	return "contacts"
}
