package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)
func main() {
 r := fasthttprouter.New()
 r.GET("/", func(ctx *fasthttp.RequestCtx) {
 Path := "./public/index.html"
 ctx.Response.Header.Set("Cache-Control", "public")
 ctx.SendFile(Path)
})
r.GET("/docs.html", func(ctx *fasthttp.RequestCtx) {
  ctx.SendFile("./public/docs.html")
})
r.GET("./help.html", func(ctx *fasthttp.RequestCtx) {
    ctx.SendFile("./public/help.html")
})
 fasthttp.ListenAndServe(":8080", r.Handler)
}

