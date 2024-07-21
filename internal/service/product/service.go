package product

import "fmt"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx >= 1 && idx <= len(allProducts) {
		return &allProducts[idx-1], nil
	}
	return nil, fmt.Errorf("wrog index: %d", idx)
}
