package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) DeleteStudent(ctx context.Context, student *models.Student) error {
	const query = `
		delete from students
		where name = $1
	`

	_, err := d.pool.Exec(ctx, query, student.Name)
	if err != nil {
		return err
	}

	return nil
}
