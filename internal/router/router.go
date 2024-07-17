package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PredictionRequest struct {
	PixelArray []int `json:"pixelArray"`
}

func sendInferRequest(requestBody []byte) ([]byte, error) {
	url := "http://localhost:3001/infer"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func printPixelsHandler(rw http.ResponseWriter, request *http.Request) {
	var req PredictionRequest

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	pixelArray := req.PixelArray
	fmt.Println(pixelArray)

	rw.WriteHeader(http.StatusOK)
}

func predictHandler(rw http.ResponseWriter, request *http.Request) {
	jsonBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error (predictHandler) when reading request body: %v\n", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := sendInferRequest(jsonBody)
	if err != nil {
		log.Printf("Error (predictHandler) when sending infer request: %v\n", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(output)
	if err != nil {
		log.Printf("Error (predictHandler) writing response in predictHandler: %v\n", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Init() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/print", printPixelsHandler)
	http.HandleFunc("/predict", predictHandler)
}

func Run(addr string) error {
	fmt.Printf("Listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}
