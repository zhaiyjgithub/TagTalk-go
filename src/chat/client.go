package chat

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zhaiyjgithub/TagTalk-go/src/model"
	"log"
	"net/http"
	"time"
)

//Client 作为中间人，用来维护App端跟server端的联系
type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan *model.Message
	uid string
}

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9)/ 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func (c *Client) readFromStream()  {
	//出现任何异常，最后登出并且释放资源
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))

		return nil
	})

	for  {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		//将全部换行符替换成空格，最后去除
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		fmt.Printf("msg: %s", string(msg))
		message := &model.Message{}
		message.RoomID = 1
		message.Text = string(msg)
		c.hub.broadcast <- message
	}
}

func (c *Client) writeToStream()  {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <- c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				//读取channel异常，关闭这次连接
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, _ = w.Write([]byte(message.PostMessage.Text))

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	param := r.URL.Query()

	uid := ""
	if param["uid"] != nil && len(param["uid"]) > 0 {
		uid = param["uid"][0]
	}else {
		log.Println("uid is need.")
		return
	}

	client := &Client{hub: hub, conn: conn, send:make(chan *model.Message), uid: uid}
	client.hub.register <- client

	go client.readFromStream()
	go client.writeToStream()
}

