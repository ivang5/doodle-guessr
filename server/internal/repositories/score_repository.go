package repositories

import (
	"log"

	"github.com/ivang5/doodle-guessr/server/internal/db"
	"github.com/ivang5/doodle-guessr/server/internal/models"
)

func InsertScore(score models.Score) (models.Score, error) {
	query := "INSERT INTO leaderboard (name, points) VALUES ($1, $2) RETURNING id"

	err := db.DB().QueryRow(query, score.Name, score.Points).Scan(&score.Id)
	if err != nil {
		log.Println("Error (InsertScore) when executing query")
		log.Printf("   |_ %v\n", err.Error())
		return score, err
	}

	return score, nil
}

func ReadScores() ([]models.Score, error) {
	query := "SELECT * FROM leaderboard"

	rows, err := db.DB().Query(query)
	if err != nil {
		log.Println("Error (ReadScores) when executing query")
		log.Printf("   |_ %v\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var scores []models.Score
	for rows.Next() {
		var score models.Score

		if err := rows.Scan(&score.Id, &score.Name, &score.Points); err != nil {
			log.Println("Error (ReadScores) when scanning row")
			log.Printf("   |_ %v\n", err.Error())
			return nil, err
		}

		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error (ReadScores) when checking db rows")
		log.Printf("   |_ %v\n", err.Error())
		return nil, err
	}

	return scores, nil
}
