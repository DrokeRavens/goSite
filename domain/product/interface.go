package product

import (
	"github.com/site/entity"
)

type Reader interface {
	Buscar(id int32) []*entity.Product
}
type Writer interface {
	Create(e *entity.Product) (int32, error)
	Update(e *entity.Product) error
	Delete(id int32) error
}
type Repository interface {
	Reader
	Writer
}
