package entity

import (
	"errors"
	"time"

	"github.com/luizgr/go-expert-api/pkg/entity"
)

var (
	ErrIDRequired      = errors.New("ID is required")
	ErrNameRequired    = errors.New("name is required")
	ErrPriceRequired   = errors.New("price is required")
	ErrIDInvalid       = errors.New("ID is invalid")
	ErrNameInvalid     = errors.New("name is invalid")
	ErrPriceInvalid    = errors.New("price is invalid")
	ErrProductNotFound = errors.New("product not found")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDInvalid
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Price <= 0 {
		return ErrPriceInvalid
	}
	return nil
}
