package repository

type Authorization interface {
}

type Service struct {
}

func NewRepository() *Service {
	return &Service{}
}