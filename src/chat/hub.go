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
					clients := getClientsByRoomID(message.RoomID, h.clients)
					for _, client := range clients {
						select {
						//这里使用select等待方式，等待向某一个client发送msg,如果发送成功就继续。
						// 如果channel已经被关闭或者有其他问题。就关闭该channel，保证内存不发生泄漏
							case client.send <- message:
						default:
							close(client.send)
							delete(h.clients, client.uid)
						}
					}
		}
	}
}

func getClientsByRoomID(roomID int64, hubClients map[string]*Client) []*Client {
	clients := make([]*Client, 0)
	for k,v := range hubClients {
		fmt.Printf("k = %s, uid = %d", k, v.uid)
		clients = append(clients, v)
	}

	return clients
}