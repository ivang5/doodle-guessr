package main

import (
	"github.com/ivang5/doodle-guessr/server/handler"
	"github.com/ivang5/doodle-guessr/server/router"
)

func main() {
	addr := "localhost:3000"

	r := router.New()
	r.Static("", "./static")

	v1 := r.Group("/api")
	h := handler.New()
	h.Register(v1)

	r.Logger.Fatal(r.Start(addr))
}
