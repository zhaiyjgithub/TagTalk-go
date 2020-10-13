package model

import (
	"time"
)

type GenderName int
type SystemStatus int

const (
	Male GenderName = 0
	Female GenderName = 1
)

const (
	Normal SystemStatus = 0
	Blocked SystemStatus = 1
)

type User struct {
	ID int
	Name string
	Avatar string
	Pin string
	Gender GenderName
	Email string
	Phone string
	LastLoginDate time.Time
	Status SystemStatus

	Token string
	TokenExpires time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserTag struct {
	ID int
	UserID int
	Name string
}

type UserImageWall struct {
	ID int
	UserID int
	Name string
}