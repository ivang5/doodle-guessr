package predict

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

func PredictHandler(rw http.ResponseWriter, request *http.Request) {
	jsonBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println("Error (PredictHandler) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := sendInferRequest(jsonBody)
	if err != nil {
		log.Println("Error (PredictHandler) when sending infer request")
		log.Printf("   |_ %v\n", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(output)
	if err != nil {
		log.Println("Error (PredictHandler) when writing response")
		log.Printf("   |_ %v\n", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
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

func PrintPixelsHandler(rw http.ResponseWriter, request *http.Request) {
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
