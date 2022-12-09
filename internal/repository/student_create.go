package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) CreateStudent(ctx context.Context, student *models.Student) (int, error) {
	const querySelectScientificDirector = `
		select tutor_id
		from tutors
		where (name || ' ' || surname) = $1	
	`
	var tutorId int
	err := d.pool.QueryRow(ctx, querySelectScientificDirector, student.ScientificDirector).Scan(&tutorId)

	const query = `
		insert into students
		(
			tutor_id, name, 
			surname, password_hash, 
			salt, course_work_subject, 
			diploma_work_subject, course
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		returning student_id;
	`

	err = d.pool.QueryRow(ctx, query, tutorId, student.Name,
		student.Surname, student.PasswordHash, student.Salt, student.CourseWorkSubject, student.DiplomaWorkSubject,
		student.Course).Scan(&student.ID)

	return student.ID, err
}
