package services

import (
	"github.com/ivang5/doodle-guessr/server/internal/models"
	"github.com/ivang5/doodle-guessr/server/internal/repositories"
)

func AddHighscoreToLeaderboard(highscore models.Highscore) (models.Highscore, error) {
	return repositories.InsertHighscore(highscore)
}

func ReadHighscoresFromLeaderboard() ([]models.Highscore, error) {
	return repositories.ReadHighscores()
}
