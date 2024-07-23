package util

import (
	"bytes"
	"io"
	"net/http"
)

func SendInferRequest(requestBody []byte) ([]byte, error) {
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

type InferRequest struct {
	Pixels []int `json:"pixels"`
}

type InferResponse struct {
	Prediction string `json:"prediction"`
}
