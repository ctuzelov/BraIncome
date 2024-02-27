package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CoursesService struct {
	repo repository.Courses
}

func NewCourseService(repository repository.Courses) *CoursesService {
	return &CoursesService{
		repo: repository,
	}
}

func (c *CoursesService) Insert(course models.Course) error {
	return c.repo.Insert(course)
}

func (c *CoursesService) Update(course models.Course) error {
	return c.repo.Update(course)
}

func (c *CoursesService) GetAll() ([]models.Course, error) {
	return c.repo.GetAll()
}

func (c *CoursesService) GetById(id primitive.ObjectID) (models.Course, error) {
	return c.repo.GetById(id)
}

func (c *CoursesService) GetByInstructor(instructor string) ([]models.Course, error) {
	return c.repo.GetByInstructor(instructor)
}

func (c *CoursesService) GetByRating() ([]models.Course, error) {
	return c.repo.GetByRating()
}

func (c *CoursesService) GetByEnrolled() ([]models.Course, error) {
	return c.repo.GetByEnrolled()
}

func (c *CoursesService) GetByCategory(category string) ([]models.Course, error) {
	return c.repo.GetByCategory(category)
}

func (c *CoursesService) GetByTitle(title string) ([]models.Course, error) {
	return c.repo.GetByTitle(title)
}
