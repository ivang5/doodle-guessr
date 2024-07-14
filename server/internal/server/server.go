package server

import (
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "doodle-guessr server is running\n")
}

func Run(addr string) error {
	http.HandleFunc("/health_check", healthCheck)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("failed to run server on %s\n", addr)
		return err
	}

	return nil
}
