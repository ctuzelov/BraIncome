package repository

import (
	"braincome/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authentication interface {
	InsertUser(models.User) error
	UserByEmail(email string) (models.User, error)
	UserByName(name string) (models.User, error)
	UserByToken(token string) (models.User, error)
	UserById(id primitive.ObjectID) (models.User, error)
	SetToken(id primitive.ObjectID, token string) error
	RemoveToken(token string) error
	SwitchRole(token string) error
}

type Courses interface {
	Insert(models.Course) error
	Update(models.Course) error
	GetAll() ([]models.Course, error)
	GetSeveral(limit int, offset int) ([]models.Course, error)
	GetById(id primitive.ObjectID) (models.Course, error)
	GetByInstructor(instructor string) ([]models.Course, error)
	GetByRating() ([]models.Course, error)
	GetByEnrolled() ([]models.Course, error)
	GetByCategory(category string) ([]models.Course, error)
	GetByTitle(title string) ([]models.Course, error)
	GetByLanguage(language string) ([]models.Course, error)
}

type Instructor interface {
	Insert(models.Instructor) error
	Update(models.Instructor) error
	GetByName(name string) (models.Instructor, error)
	GetByEmail(email string) (models.Instructor, error)
}

type Repository struct {
	Authentication
	Courses
	Instructor
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
		Courses:        NewCoursesMongo(db),
		Instructor:     NewInstructorMongo(db),
	}
}
