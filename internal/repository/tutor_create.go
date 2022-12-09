package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) CreateTutor(ctx context.Context, tutor *models.Tutor) (int, error) {
	const query = `
		insert into tutors
		(
			name, surname, 
			password_hash, salt, 
			email, phone,
			has_free_places
		)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning student_id;
	`

	err := d.pool.QueryRow(ctx, query, tutor.Name, tutor.Surname,
		tutor.PasswordHash, tutor.Salt, tutor.Email, tutor.Phone, tutor.IsFreePlaces).Scan(&tutor.ID)

	return tutor.ID, err
}
