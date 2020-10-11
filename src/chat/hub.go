package chat

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}

func (h *Hub) Run()  {
	for {
		select {
		case client := <- h.register:
			h.clients[client] = true

			case client := <- h.unregister:
				delete(h.clients, client)
				close(client.send)

				case message := <- h.broadcast:
					for client := range h.clients {
						select { //这里使用select等待方式，等待向某一个client发送msg,如果发送成功就继续。
						// 如果channel已经被关闭或者有其他问题。就关闭该channel，保证内存不发生泄漏
							case client.send <- message:
						default:
							close(client.send)
							delete(h.clients, client)
						}
					}
		}
	}
}