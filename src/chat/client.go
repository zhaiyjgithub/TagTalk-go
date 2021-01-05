package chat

import (
	"encoding/json"
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
	pingPeriod = (pongWait * 9)/ 10 //接近timeout，就发送一个ping message到App
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
		//这里会将客户端传递上来的message string 转成 message model
		m := &model.Message{}
		err = json.Unmarshal(msg, &m)
		if err != nil {
			return
		}

		c.hub.broadcast <- m
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

			msg, err := json.Marshal(message)
			if err != nil {
				fmt.Printf("marshal message error: %s", err.Error())
			}
			_, _ = w.Write(msg)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
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

	chatID := ""
	if param["chatID"] != nil && len(param["chatID"]) > 0 {
		chatID = param["chatID"][0]
	}else {
		log.Println("uid is need.")
		return
	}

	fmt.Printf("Login chatID: %s \n", chatID)
	client := &Client{hub: hub, conn: conn, send:make(chan *model.Message, 512), uid: chatID}
	client.hub.register <- client

	go client.readFromStream()
	go client.writeToStream()
}

