package chat

import (
	"fmt"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
)

type Hub struct {
	clients map[string]*Client
	broadcast chan *model.Message
	register chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
		broadcast: make(chan *model.Message),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run()  {
	for {
		select {
		case client := <- h.register:
			h.clients[client.uid] = client
		case client := <- h.unregister:
			delete(h.clients, client.uid)
			close(client.send)
		case message := <- h.broadcast:
			go h.serveBroadcast(message)
		}
	}
}

func (h *Hub)serveBroadcast(message *model.Message)  {
	client := getSingleClient(h.clients, message.ChannelID)
	handleSingleClientMessage(client, message)
}

func handleSingleClientMessage(client *Client, message *model.Message)  {
	if client != nil {
		client.send <- message
	}
}

func handleMultiClientMessage(hubClients map[int64]*Client, uid int64)  {
	//clients := getRoomClients(message.RoomID, h.clients)
	//for _, client := range clients {
	//	select {
	//	//这里使用select等待方式，等待向某一个client发送msg,如果发送成功就继续。
	//	// 如果channel已经被关闭或者有其他问题。就关闭该channel，保证内存不发生泄漏
	//	case client.send <- message:
	//	default:
	//		close(client.send)
	//		delete(h.clients, client.uid)
	//	}
	//}
}

func getSingleClient(hubClients map[string]*Client, uid string) *Client  {
	return hubClients[uid]
}

func getRoomClients(hubClients map[int64]*Client, channelID int64) []*Client {
	clients := make([]*Client, 0)
	for k,v := range hubClients {
		fmt.Printf("k = %s, uid = %d", k, v.uid)
		clients = append(clients, v)
	}

	return clients
}