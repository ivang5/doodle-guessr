package router

import (
	"fmt"
	"net/http"

	"github.com/ivang5/doodle-guessr/server/internal/predict"
)

func Init() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/predict", predict.PredictHandler)
	http.HandleFunc("/print", predict.PrintPixelsHandler)
}

func Run(addr string) error {
	fmt.Printf("Listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}
