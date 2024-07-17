package main

import (
	"fmt"

	"github.com/ivang5/doodle-guessr/server/internal/router"
)

func main() {
	addr := "localhost:6969"

	router.Init()
	if err := router.Run(addr); err != nil {
		fmt.Printf("failed to run server on %s\n", addr)
	}
}
