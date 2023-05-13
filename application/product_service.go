package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64, quantity int) (ProductInterface, error) {

	input := NewProductDto{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	product, err := NewProduct(input)

	if err != nil {
		return nil, err
	}

	saved_product, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return saved_product, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return nil, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	product.Disable()

	result, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}
