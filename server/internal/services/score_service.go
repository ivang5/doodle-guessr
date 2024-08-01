package services

import (
	"github.com/ivang5/doodle-guessr/server/internal/models"
	"github.com/ivang5/doodle-guessr/server/internal/repositories"
)

func AddScoreToLeaderboard(score models.Score) (models.Score, error) {
	return repositories.InsertScore(score)
}

func ReadScoresFromLeaderboard() ([]models.Score, error) {
	return repositories.ReadScores()
}
