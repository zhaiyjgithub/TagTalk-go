package model

import "time"

type MessageType int
type ChannelType int

const (
	Text MessageType = 0
	Image MessageType = 1
	Audio MessageType = 2
	Video MessageType = 3
	WebView MessageType = 4
	MapView MessageType = 5
)

const (
	Single ChannelType = 1
	Group ChannelType = 2
	Broadcast ChannelType = 3
)

type mapView struct {
	lat float64
	lng float64
}

type Message struct {
	ID string `json:"id"`

	NickName string `json:"nickName"`
	Avatar string	`json:"avatar"`
	SenderId string	`json:"senderId"`
	Mentions []string `json:"mentions"`

	ChannelType ChannelType 	`json:"channelType"`
	ChannelID string `json:"channelId"`

	MessageType MessageType `json:"messageType"`
	Message string `json:"message"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}