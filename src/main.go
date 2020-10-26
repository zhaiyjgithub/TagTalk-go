package main

import (
	"github.com/kataras/iris/v12"
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
	app := iris.New()

	hub := chat.NewHub()
	go hub.Run()
	app.Any("/ws", func(ctx iris.Context) {
		chat.ServeWs(hub, ctx.ResponseWriter(), ctx.Request())
	})

	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M

}