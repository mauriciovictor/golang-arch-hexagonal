package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (p ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct(name, price)

	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	result, err := p.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s ProductService) Enable(produc ProductInterface) (ProductInterface, error) {
	err := produc.Enable()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(produc)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s ProductService) Disable(produc ProductInterface) (ProductInterface, error) {
	err := produc.Disable()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(produc)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}
