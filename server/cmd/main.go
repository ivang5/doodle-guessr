package main

import (
	"fmt"

	"github.com/ivang5/doodle-guessr/server/internal/server"
)

func main() {
	addr := "localhost:6969"

	fmt.Printf("Listening on %s\n", addr)
	server.Init()
	if err := server.ListenAndServe(addr); err != nil {
		fmt.Printf("failed to run server on %s\n", addr)
	}
}
