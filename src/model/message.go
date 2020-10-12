package model

import "time"

type MessageType int

const (
	Text MessageType = 0
	Image MessageType = 1
	Audio MessageType = 2
	Video MessageType = 3
	WebView MessageType = 4
	MapView MessageType = 5
)

type mapView struct {
	lat float64
	lng float64
}

type Message struct {
	ID string
	RoomID int64
	User *User
	Mentions []*User
	CreatedAt *time.Time
	UpdatedAt *time.Time
	PostMessage
}

type PostMessage struct {
	RoomID int64
	NickName string
	Avatar string
	Type MessageType

	Text string
	ImageURL string
	AudioURL string
	VideoURL string
	WebViewURL string
	MapView mapView
}
