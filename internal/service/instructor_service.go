package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"
)

type InstructorService struct {
	repo repository.Instructor
}

func NewInstructorService(repository repository.Instructor) *InstructorService {
	return &InstructorService{
		repo: repository,
	}
}

func (c *InstructorService) Insert(instructor models.Instructor) error {
	return c.repo.Insert(instructor)
}

func (c *InstructorService) Update(instructor models.Instructor) error {
	return c.repo.Update(instructor)
}

func (c *InstructorService) GetByName(name string) (models.Instructor, error) {
	return c.repo.GetByName(name)
}

func (c *InstructorService) GetByEmail(email string) (models.Instructor, error) {
	return c.repo.GetByEmail(email)
}
