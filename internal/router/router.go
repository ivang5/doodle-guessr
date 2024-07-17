package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type PredictionRequest struct {
	PixelArray []int `json:"pixelArray"`
}

func runInferCommand(pixelArrayJson string) (string, error) {
	// modelPath := "model/model.py"
	modelPath := "model/test.py"
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		return "", err
	}

	cmd := exec.Command("python3", modelPath, pixelArrayJson)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := out.String()

	return result, nil
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
		log.Printf("Error occured in predictHandler: %v\n", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonString := string(jsonBody)

	output, err := runInferCommand(jsonString)
	if err != nil {
		log.Printf("Error occured in predictHandler: %v\n", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(output)

	rw.WriteHeader(http.StatusOK)
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
