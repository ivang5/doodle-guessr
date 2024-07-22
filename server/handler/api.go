package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Predict(c echo.Context) error {
	var req predictRequest
	var resp predictResponse

	if err := c.Bind(&req); err != nil {
		log.Println("Error (PredictHandler) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		log.Println("Error (PredictHandler) when marshalling request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	responseBody, err := sendInferRequest(jsonBody)
	if err != nil {
		log.Println("Error (PredictHandler) when sending infer request")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err = json.Unmarshal(responseBody, &resp); err != nil {
		log.Println("Error (PredictHandler) when unmarshalling infer response")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp.Prediction)
}

// TODO: Mire
// adapt sending request to echo if needed
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

func (h *Handler) Print(c echo.Context) error {
	var req predictRequest
	if err := c.Bind(&req); err != nil {
		log.Println("Error (PredictHandler) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Println(req.Pixels)

	return c.NoContent(http.StatusOK)
}
