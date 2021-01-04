package model

import "time"

type GenderType int

const (
	Male GenderType = 0
	Female GenderType = 1
)

type User struct {
	ID         int             	`gorm:"column:id;primary_key" json:"-"`
	Name       string  			`gorm:"column:name"`
	Email      string 			`gorm:"column:email"`
	Phone      string 			`gorm:"column:phone"`
	Bio        string 			`gorm:"column:bio"`
	Avatar	   string 		    `gorm:"column:avatar"`
	Level      int  			`gorm:"column:level"`
	Password   string 			`grom:"column:password"`
	Gender     GenderType 		`gorm:"column:gender"`
	ChatID     string 			`gorm:"column:chat_id"`
	CreatedAt  time.Time      	`gorm:"column:created_at" json:"-"`
	UpdatedAt  time.Time      	`gorm:"column:updated_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
