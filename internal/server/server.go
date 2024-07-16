package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PredictionRequest struct {
	PixelArray []int `json:"pixelArray"`
}

func predictHandler(rw http.ResponseWriter, request *http.Request) {
	var req PredictionRequest

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	pixelArray := req.PixelArray

	fmt.Println(pixelArray)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{"status": "success"})
}

func Init() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/predict", predictHandler)
}

func ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
