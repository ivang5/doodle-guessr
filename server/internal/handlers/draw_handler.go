package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ivang5/doodle-guessr/server/internal/util"
	"github.com/labstack/echo/v4"
)

func Predict(c echo.Context) error {
	var req util.InferRequest
	var resp util.InferResponse

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

	responseBody, err := util.SendInferRequest(jsonBody)
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

func Print(c echo.Context) error {
	var req predictRequest
	if err := c.Bind(&req); err != nil {
		log.Println("Error (PredictHandler) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Println(req.Pixels)

	return c.NoContent(http.StatusOK)
}

type predictRequest struct {
	Pixels []int `json:"pixels"`
}

type predictResponse struct {
	Prediction string `json:"prediction"`
}
