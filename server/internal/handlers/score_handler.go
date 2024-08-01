package handlers

import (
	"log"
	"net/http"

	"github.com/ivang5/doodle-guessr/server/internal/models"
	"github.com/ivang5/doodle-guessr/server/internal/services"
	"github.com/ivang5/doodle-guessr/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func SetScore(c echo.Context) error {
	var req SetScoreRequest

	if err := c.Bind(&req); err != nil {
		log.Println("Error (SetScore) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorAsMap(err))
	}

	score := models.Score{
		Name:   req.Name,
		Points: req.Points,
	}

	score, err := services.AddScoreToLeaderboard(score)
	if err != nil {
		log.Println("Error (SetScore) when adding score to leaderboard")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
	}

	resp := InsertScoreResponse{
		Id:     int(score.Id),
		Name:   score.Name,
		Points: score.Points,
	}

	return c.JSON(http.StatusOK, resp)
}

func ReadScores(c echo.Context) error {
	scores, err := services.ReadScoresFromLeaderboard()
	if err != nil {
		log.Println("Error (ReadScores) when reading scores from leaderboard")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
	}

	var respScores []Score
	for _, score := range scores {
		respScores = append(respScores, Score{
			Name:   score.Name,
			Points: score.Points,
		})
	}

	resp := ReadScoresResponse{
		Scores: respScores,
	}

	return c.JSON(http.StatusOK, resp)
}

type SetScoreRequest struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type InsertScoreResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type Score struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type ReadScoresResponse struct {
	Scores []Score `json:"scores"`
}
