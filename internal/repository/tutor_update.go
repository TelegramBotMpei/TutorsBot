package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) UpdateTutor(ctx context.Context, tutor *models.Tutor) error {
	const query = `
		update tutor
		set name = $2,
		surname = $3,
		email = $4,
		phone = $5,
		has_free_places = #6
		where student_id = $1;
	`
	_, err := d.pool.Exec(ctx, query, tutor.ID, tutor.Name, tutor.Surname, tutor.Email, tutor.Phone, tutor.IsFreePlaces)

	if err != nil {
		return err
	}

	return nil
}
