package main

import (
	"github.com/zhaiyjgithub/TagTalk-go/src/chat"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.URL)
	//if r.URL.Path != "/" {
	//	http.Error(w, "Not found", http.StatusNotFound)
	//	return
	//}
	//if r.Method != "GET" {
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//	return
	//}
	http.ServeFile(w, r, "./src/chat/home.html")
}

func main() {
	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}