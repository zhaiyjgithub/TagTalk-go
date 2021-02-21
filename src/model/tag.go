package model

type Tag struct {
	ID     int   `gorm:"column:id;primary_key" json:"-"`
	ChatID string `gorm:"column:chat_id"`
	Names string `gorm:"column:names"`
}

// TableName sets the insert table name for this struct type
func (c *Tag) TableName() string {
	return "tags"
}
