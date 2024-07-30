package handlers

import (
	"log"
	"net/http"

	"github.com/ivang5/doodle-guessr/server/internal/models"
	"github.com/ivang5/doodle-guessr/server/internal/services"
	"github.com/ivang5/doodle-guessr/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func SetHighscore(c echo.Context) error {
	var req SetHighscoreRequest

	if err := c.Bind(&req); err != nil {
		log.Println("Error (SetHighscore) when reading request body")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorAsMap(err))
	}

	highscore := models.Highscore{
		Name:  req.Name,
		Score: req.Score,
	}

	exists, err := services.NameExistsInLeaderboard(highscore.Name)
	if err != nil {
		log.Println("Error (SetHighscore) when checking for name in leaderboard")
		log.Printf("   |_ %v\n", err.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
	}

	if exists {
		err := services.UpdateHighscoreInLeaderboard(highscore)
		if err != nil {
			log.Println("Error (SetHighscore) when updating highscore in leaderboard")
			log.Printf("   |_ %v\n", err.Error())
			return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
		}

		resp := UpdateHighscoreResponse{
			Name:  highscore.Name,
			Score: highscore.Score,
		}

		return c.JSON(http.StatusOK, resp)
	} else {
		highscore, err := services.AddHighscoreToLeaderboard(highscore)

		if err != nil {
			log.Println("Error (SetHighscore) when adding highscore to leaderboard")
			log.Printf("   |_ %v\n", err.Error())
			return c.JSON(http.StatusInternalServerError, utils.ErrorAsMap(err))
		}

		resp := InsertHighscoreResponse{
			Id:    int(highscore.Id),
			Name:  highscore.Name,
			Score: highscore.Score,
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func ReadHighscores(c echo.Context) error {
	return nil
}

type SetHighscoreRequest struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type InsertHighscoreResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type UpdateHighscoreResponse struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
