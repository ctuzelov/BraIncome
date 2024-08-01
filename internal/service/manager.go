package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	SignUp(models.User) error
	SignIn(login, password string) (models.User, error)
	GetByToken(token string) (models.User, error)
	LogOut(token string) error
	MakeAdmin(email string) error
	DeleteTokensByEmail(email string) error
}

type Courses interface {
	Insert(course models.Course) error
	Update(course models.Course) error
	GetAll() ([]models.Course, error)
	GetSeveral(limit int, offset int) ([]models.Course, error)
	GetById(id primitive.ObjectID) (models.Course, error)
	GetByInstructor(instructor string) ([]models.Course, error)
	GetByRating() ([]models.Course, error)
	GetByEnrolled() ([]models.Course, error)
	GetByCategory(category string) ([]models.Course, error)
	GetByTitle(title string) ([]models.Course, error)
}

type Instructor interface {
	Insert(course models.Instructor) error
	Update(course models.Instructor) error
	GetByName(name string) (models.Instructor, error)
	GetByEmail(email string) (models.Instructor, error)
}

type Service struct {
	User
	Courses
	Instructor
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:       NewAuthService(repo.Authentication),
		Courses:    NewCourseService(repo.Courses),
		Instructor: NewInstructorService(repo.Instructor),
	}
}
