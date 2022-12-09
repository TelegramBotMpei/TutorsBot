package repository

import (
	"context"
	"github.com/TelegramBotMPEI/internal/models"
)

type StudentRepo interface {
	CreateStudent(ctx context.Context, student *models.Student, password string) (int, error)
	ReadStudent(ctx context.Context, id int) (*models.Student, error)
	UpdateStudent(ctx context.Context, student *models.Student) error
	DeleteStudent(ctx context.Context, student *models.Student) error
}

type TutorRepo interface {
	CreateTutor(ctx context.Context, tutor *models.Tutor, password string) (int, error)
	ReadTutor(ctx context.Context, id int) (*models.Tutor, error)
	UpdateTutor(ctx context.Context, tutor *models.Tutor) error
	DeleteTutor(ctx context.Context, tutor *models.Tutor) error
	GetScientificInterests(ctx context.Context, tutor *models.Tutor) ([]string, error)
}
