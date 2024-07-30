package repositories

import (
	"log"

	"github.com/ivang5/doodle-guessr/server/internal/db"
	"github.com/ivang5/doodle-guessr/server/internal/models"
)

func InsertHighscore(highscore models.Highscore) (models.Highscore, error) {
	query := "INSERT INTO highscores (name, score) VALUES ($1, $2) RETURNING id"

	err := db.DB().QueryRow(query, highscore.Name, highscore.Score).Scan(&highscore.Id)
	if err != nil {
		log.Println("Error (InsertHighscore) when executing query")
		log.Printf("   |_ %v\n", err.Error())
		return highscore, err
	}

	return highscore, nil
}

func ReadHighscores() ([]models.Highscore, error) {
	query := "SELECT * FROM highscores"

	rows, err := db.DB().Query(query)
	if err != nil {
		log.Println("Error (ReadHighscores) when executing query")
		log.Printf("   |_ %v\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var highscores []models.Highscore
	for rows.Next() {
		var highscore models.Highscore

		if err := rows.Scan(&highscore.Id, &highscore.Name, &highscore.Score); err != nil {
			log.Println("Error (ReadHighscores) when scanning row")
			log.Printf("   |_ %v\n", err.Error())
			return nil, err
		}

		highscores = append(highscores, highscore)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error (ReadHighscores) when checking db rows")
		log.Printf("   |_ %v\n", err.Error())
		return nil, err
	}

	return highscores, nil
}
