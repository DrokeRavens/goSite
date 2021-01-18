package product

import "github.com/site/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Buscar(id int32) []*entity.Product {
	return s.repo.Buscar(id)
}
