package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() error
	Enable() error
	Disable()
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	GetQuantity() int
	ChangePrice(price float64) error
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64, quantity int) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID       string  `valid:"uuidv4"`
	Name     string  `valid:"required"`
	Status   string  `valid:"required"`
	Price    float64 `valid:"float,optional"`
	Quantity int     `valid:"int,optional"`
}

func (p *Product) IsValid() error {
	if p.Name == "" {
		return errors.New("Invalid name")
	} else if p.Price < 0 {
		return errors.New("Invalid price")
	} else if p.Status != DISABLED && p.Status != ENABLED {
		return errors.New("Invalid status")
	} else if p.Quantity < 0 {
		return errors.New("Invalid quantity")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Enable() error {
	err := p.IsValid()

	if err != nil {
		return err
	}

	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() {
	p.Status = DISABLED
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetQuantity() int {
	return p.Quantity
}

func (p *Product) ChangePrice(price float64) error {

	if price < 0 {
		return errors.New("Invalid price")
	}

	p.Price = price

	err := p.IsValid()

	if err != nil {
		return err
	}
	return nil
}

func (p *Product) ChangeQuantity(quantity int) error {

	p.Quantity = quantity

	err := p.IsValid()

	if err != nil {
		return err
	}
	return nil
}

type NewProductDto struct {
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(input NewProductDto) (*Product, error) {
	product := Product{
		Name:     input.Name,
		Status:   DISABLED,
		Price:    input.Price,
		Quantity: input.Quantity,
		ID:       uuid.NewString(),
	}

	err := product.IsValid()

	if err != nil {
		return nil, err
	}

	return &product, nil
}
