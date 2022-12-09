package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) ReadTutor(ctx context.Context, id int) (*models.Tutor, error) {
	const query = `
		select
		name, surname,
		email, phone,
		has_free_places
		where id = $1
	`

	row, err := d.pool.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	student := models.Tutor{
		ID: id,
	}

	err = row.Scan(&student.Name, &student.Surname, &student.Email,
		&student.Phone, &student.IsFreePlaces)

	if err != nil {
		return nil, err
	}

	return &student, nil
}
