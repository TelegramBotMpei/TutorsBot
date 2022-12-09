package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) DeleteTutor(ctx context.Context, tutor *models.Tutor) error {
	const query = `
		delete from students
		where name = $1
	`

	_, err := d.pool.Exec(ctx, query, tutor.Name)
	if err != nil {
		return err
	}

	return nil
}
