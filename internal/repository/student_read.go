package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) ReadStudent(ctx context.Context, id int) (*models.Student, error) {
	const query = `
		select tutor_id, 
		name, surname,
		course_work_subject, diploma_work_subject,
		course
		where id = $1
	`

	row, err := d.pool.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	student := models.Student{
		ID: id,
	}

	err = row.Scan(&student.TutorID, &student.Name, &student.Surname, &student.CourseWorkSubject,
		&student.DiplomaWorkSubject, &student.Course)

	if err != nil {
		return nil, err
	}

	return &student, nil
}
