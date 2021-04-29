package athlete

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AthletesRepository struct {
	db *pgxpool.Pool
}

const insertAthleteSQL = "INSERT INTO athletes (name, category) VALUES (:name,:category) RETURNING id"

func (u *AthletesRepository) CreateUser(athlete Athlete) (*Athlete, error) {
	var err error

	if err != nil {
		return nil, err
	}

	var athleteId string
	if err = u.db.QueryRow(context.Background(), insertAthleteSQL, athlete.Name, athlete.Category).Scan(&athleteId); err != nil {
		return nil, errors.New("error creating the athlete")
	}
	athlete.ID = athleteId
	return &athlete, nil
}

func CreateRepository(db *pgxpool.Pool) *AthletesRepository {
	return &AthletesRepository{db: db}
}
