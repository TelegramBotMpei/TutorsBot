package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

func (d *postgresDB) UpdateStudent(ctx context.Context, student *models.Student) error {
	const query = `
		update students
		set tutor_id = $2,
		name = $3, 
		surname = #4, 
		course_work_subject = $5, 
    	diploma_work_subject = $6, 
		course = $7
		where student_id = $1;
	`
	_, err := d.pool.Exec(ctx, query, student.ID, student.TutorID, student.Name, student.Surname,
		student.CourseWorkSubject, student.DiplomaWorkSubject, student.Course)

	if err != nil {
		return err
	}

	return nil
}
