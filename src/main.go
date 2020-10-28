package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zhaiyjgithub/TagTalk-go/src/chat"
	"github.com/zhaiyjgithub/TagTalk-go/src/controller"
	"github.com/zhaiyjgithub/TagTalk-go/src/service"
	"github.com/zhaiyjgithub/TagTalk-go/src/utils"
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

	userParty := app.Party(utils.APIUser)
	mvc.Configure(userParty, userMVC)

	chatParty := app.Party(utils.APIChat)
	mvc.Configure(chatParty, chatMVC)

	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M
}

func chatMVC(app *mvc.Application)  {
	app.Handle(new(controller.ChatController))
}

func userMVC(app *mvc.Application)  {
	userService := service.NewUserService()

	app.Register(userService)
	app.Handle(new(controller.UserController))
}