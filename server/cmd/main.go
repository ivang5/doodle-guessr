package main

import (
	"fmt"

	"github.com/ivang5/doodle-guessr/server/internal/router"
)

func main() {
	addr := "localhost:3000"

	router.Init()
	if err := router.Run(addr); err != nil {
		fmt.Printf("Failed to run server on %s\n", addr)
		fmt.Printf("   |_ %v\n", err.Error())
	}
}
