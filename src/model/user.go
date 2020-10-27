package model

import (
	"time"
)

type GenderType int
type SystemStatus int

const (
	Male GenderType = 0
	Female GenderType = 1
)

const (
	Normal SystemStatus = 0
	Blocked SystemStatus = 1
)


type User struct {
	ID          int            `gorm:"column:id;primary_key"`
	Name        string `gorm:"column:name"`
	Phone       string `gorm:"column:phone"`
	Email       string `gorm:"column:email"`
	Gender      GenderType `gorm:"column:gender"`
	HeaderIcon  string `gorm:"column:header_icon"`
	TagListID   int  `gorm:"column:tag_list_id"`
	Description string `gorm:"column:description"`
	ImageListID int  `gorm:"column:image_list_id"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
