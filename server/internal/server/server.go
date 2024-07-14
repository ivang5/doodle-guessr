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

func Init() {
	http.HandleFunc("/health_check", healthCheck)
}

func ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
