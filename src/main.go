package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zhaiyjgithub/TagTalk-go/src/chat"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./src/chat/home.html")
}

func main() {
	app := iris.New()
	app.Post("/upload", func(ctx iris.Context) {
		//_, h, _ := ctx.FormFile("test")
		maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

		err := ctx.Request().ParseMultipartForm(maxSize)
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		//multiForm := ctx.Request().MultipartForm
		//fmt.Print(multiForm.File)
		//h1 := ctx.Request().FormValue("my")
		//fmt.Print(h1)
		//ctx.SaveFormFile(h, "./text.txt")
		_, _ = ctx.WriteString("success")

		//multiForm := ctx.Request().MultipartForm
		//file := multiForm.File

		//for _, fh := range file {
		//	ctx.SaveFormFile(fh, "./")
		//}

		_, _ = ctx.WriteString("success")
	})

	hub := chat.NewHub()
	go hub.Run()
	app.Any("/ws", func(ctx iris.Context) {
		chat.ServeWs(hub, ctx.ResponseWriter(), ctx.Request())
	})

	_ = app.Run(iris.Addr(":8090"), iris.WithPostMaxMemory(32<<20)) //max = 32M

}