package model

type Wall struct {
	ID     int   `gorm:"column:id;primary_key" json:"-"`
	ChatID string `gorm:"column:chat_id"`
	Names string `gorm:"column:names"`
}

// TableName sets the insert table name for this struct type
func (c *Wall) TableName() string {
	return "walls"
}