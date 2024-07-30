package services

import (
	"github.com/ivang5/doodle-guessr/server/internal/models"
	"github.com/ivang5/doodle-guessr/server/internal/repositories"
)

func AddHighscoreToLeaderboard(highscore models.Highscore) (models.Highscore, error) {
	return repositories.InsertHighscore(highscore)
}

func UpdateHighscoreInLeaderboard(highscore models.Highscore) error {
	return repositories.UpdateHighscore(highscore)
}

func NameExistsInLeaderboard(name string) (bool, error) {
	highscores, err := repositories.ReadHighscores()

	if err != nil {
		return false, err
	}

	for _, highscore := range highscores {
		if highscore.Name == name {
			return true, nil
		}
	}

	return false, nil
}
