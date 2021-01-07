package model

import "time"

type MessageMediaType int
type ChannelType int
type MessageCategory int

const (
	Text MessageMediaType = 0
	Image MessageMediaType = 1
	Audio MessageMediaType = 2
	Video MessageMediaType = 3
	WebView MessageMediaType = 4
	MapView MessageMediaType = 5
)

const (
	Normal MessageCategory = 0
	NewDialog MessageCategory = 1
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
	
	Category MessageCategory `json:"category"`
	MediaType MessageMediaType `json:"mediaType"`
	Message string `json:"message"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}