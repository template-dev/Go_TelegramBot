package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetProducts() []Product {
	return allProducts
}
