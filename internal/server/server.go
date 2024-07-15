package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PredictionRequest struct {
	PixelArray []int `json:"pixelArray"`
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "doodle-guessr server is running\n")
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the struct to hold the request body.
	var req PredictionRequest

	// Parse the JSON body into the struct.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Access the pixelArray from the struct.
	pixelArray := req.PixelArray

	// Your logic here, for example, logging the pixelArray or processing it.
	fmt.Println(pixelArray)

	// Respond to the client.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func Init() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/predict", predictHandler)
}

func ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
