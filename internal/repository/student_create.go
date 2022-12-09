package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
	"github.com/TelegramBotMPEI/internal/utils"
)

func (d *postgresDB) CreateStudent(ctx context.Context, student *models.Student, password string) (int, error) {
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

	salt := utils.GenerateRandomSalt(10) // TODO: MAYBE its configure parametr?

	passwordHash, err := utils.HashPassword(password + salt)
	if err != nil {
		return -1, err
	}

	err = d.pool.QueryRow(ctx, query, tutorId, student.Name,
		student.Surname, passwordHash, salt, student.CourseWorkSubject, student.DiplomaWorkSubject,
		student.Course).Scan(&student.ID)

	return student.ID, err
}
