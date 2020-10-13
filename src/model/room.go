package model

import "time"

type Room struct {
	ID int
	Channel string
	Name string
	Description string
	OwnerID int
	ManagerID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Pin string
	TagsID int
}

type Member struct {
	ID int
	RoomID int
	UserID int
}
