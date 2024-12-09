package service

import (
	"errors"
	"main/internal/repository"
	"regexp"
	"strconv"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetById(id int) (*repository.Product, bool) {
	product, ok := s.repo.GetById(id)

	return product, ok
}

func (s *ProductService) GetAll() *[]*repository.Product {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductsFiltered(priceGt float32) *[]*repository.Product {
	return s.repo.SearchByPrice(priceGt)
}

func (s *ProductService) Delete(product *repository.Product) error {
	return s.repo.Delete(product)
}

func (s *ProductService) Create(product *repository.Product) error {
	if err := s.validateProduct(product); err != nil {
		return err
	}
	return s.repo.Create(product)
}

func (s *ProductService) Patch(product *repository.Product) error {
	// Validar o produto atualizado
	if err := s.validateProduct(product); err != nil {
		return err
	}
	return s.repo.Update(product)
}

func (s *ProductService) Put(product *repository.Product) (bool, error) {
	// Validação dos campos
	if err := s.validateProduct(product); err != nil {
		return false, err
	}
	return s.repo.Put(product)
}

func (s *ProductService) validateProduct(product *repository.Product) error {
	if product.Name == "" {
		return errors.New("O campo 'name' não pode ser vazio")
	}
	if product.Quantity == 0 {
		return errors.New("O campo 'quantity' não pode ser zero")
	}
	if product.CodeValue == "" {
		return errors.New("O campo 'code_value' não pode ser vazio")
	}
	if s.repo.ExistsByCodeValue(product) {
		return errors.New("O campo 'code_value' já existe")
	}
	if product.Expiration == "" {
		return errors.New("O campo 'expiration' não pode ser vazio")
	}
	if product.Price == 0 {
		return errors.New("O campo 'price' não pode ser zero")
	}
	// Validar formato da data
	dateRegex := regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})$`)
	matches := dateRegex.FindStringSubmatch(product.Expiration)
	if matches == nil {
		return errors.New("O campo 'expiration' deve estar no formato DD/MM/AAAA")
	}

	day, err := strconv.Atoi(matches[1])
	if err != nil || day < 1 || day > 31 {
		return errors.New("Dia inválido na data de validade")
	}

	month, err := strconv.Atoi(matches[2])
	if err != nil || month < 1 || month > 12 {
		return errors.New("Mês inválido na data de validade")
	}

	year, err := strconv.Atoi(matches[3])
	if err != nil || year < 1 {
		return errors.New("Ano inválido na data de validade")
	}

	return nil
}
