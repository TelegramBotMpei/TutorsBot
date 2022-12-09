package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
	"github.com/TelegramBotMPEI/internal/utils"
)

func (d *postgresDB) CreateTutor(ctx context.Context, tutor *models.Tutor, password string) (int, error) {
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

	salt := utils.GenerateRandomSalt(10) // TODO: MAYBE its configure parametr?

	passwordHash, err := utils.HashPassword(password + salt)
	if err != nil {
		return -1, err
	}

	err = d.pool.QueryRow(ctx, query, tutor.Name, tutor.Surname,
		passwordHash, salt, tutor.Email, tutor.Phone, tutor.IsFreePlaces).Scan(&tutor.ID)

	return tutor.ID, err
}
