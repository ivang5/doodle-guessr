package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/ivang5/doodle-guessr/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func Predict(c echo.Context) error {
	var req PredictRequest
	var resp PredictResponse

	if err := c.Bind(&req); err != nil {
		log.Println("Error (Predict) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorAsMap(err))
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		log.Println("Error (Predict) when marshalling request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
	}

	responseBody, err := SendPredictRequest(jsonBody)
	if err != nil {
		log.Println("Error (Predict) when sending predict request")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
	}

	if err = json.Unmarshal(responseBody, &resp); err != nil {
		log.Println("Error (Predict) when unmarshalling predict response")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
	}

	return c.JSON(http.StatusOK, resp)
}

func SendPredictRequest(requestBody []byte) ([]byte, error) {
	url := "http://localhost:3001/predict"

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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return responseBody, nil
}

type PredictRequest struct {
	Pixels []int `json:"pixels"`
}

type PredictResponse struct {
	Prediction string  `json:"prediction"`
	Certainty  float32 `json:"certainty"`
}
