package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) GetScientificInterests(ctx context.Context, tutor *models.Tutor) ([]string, error) {
	const querySelectTutorID = `
		select tutor_id
		from tutors
		name = $1
	`

	var tutorId int
	err := d.pool.QueryRow(ctx, querySelectTutorID, tutor.Name).Scan(&tutorId)

	if err != nil {
		return []string{}, nil
	}

	const query = `
		select subject
		from scientific_interests
		where tutor_id = $1
	`

	rows, err := d.pool.Query(ctx, query, tutorId)

	if err != nil {
		return []string{}, err
	}

	scientificInterests := make([]string, 0)

	for rows.Next() {
		var scientificInterest string
		_ = rows.Scan(&scientificInterest)
		scientificInterests = append(scientificInterests, scientificInterest)
	}

	return scientificInterests, nil
}
