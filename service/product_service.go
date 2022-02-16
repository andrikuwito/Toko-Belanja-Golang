package service

import (
	"errors"
	"tokobelanja-golang/model/entity"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/repository"
)

type ProductService interface {
	CreateProduct(input input.CreateProduct) (entity.Product, error)
	GetProductByID(ID int) (entity.Product, error)
	GetProduct() ([]entity.Product, error)
	UpdateProduct(ID int, input input.UpdateProduct) (entity.Product, error)
	DeleteProduct(ID int) (entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(input input.CreateProduct) (entity.Product, error) {
	newProduct := entity.Product{}
	newProduct.Title = input.Title
	newProduct.Price = input.Price
	newProduct.Stock = input.Stock
	newProduct.CategoryID = input.CategoryID

	created, err := s.productRepository.Save(newProduct)

	if err != nil {
		return entity.Product{}, err
	}

	return created, nil
}

func (s *productService) GetProductByID(ID int) (entity.Product, error) {
	product, err := s.productRepository.GetProductByID(ID)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productService) GetProduct() ([]entity.Product, error) {
	product, err := s.productRepository.GetAllProduct()

	if err != nil {
		return product, err
	}

	return product, nil
}
func (s *productService) UpdateProduct(ID int, input input.UpdateProduct) (entity.Product, error) {
	Result, err := s.productRepository.GetProductByID(ID)

	if err != nil {
		return entity.Product{}, err
	}

	if Result.ID == 0 {
		return entity.Product{}, errors.New("not found!")
	}

	updated := entity.Product{
		ID:    ID,
		Title: input.Title,
		Price: input.Price,
		Stock: input.Stock,
	}

	productUpdate, err := s.productRepository.Update(updated)

	if err != nil {
		return productUpdate, err
	}

	return productUpdate, nil
}

func (s *productService) DeleteProduct(ID int) (entity.Product, error) {
	product, err := s.productRepository.GetProductByID(ID)

	if err != nil {
		return entity.Product{}, err
	}

	if product.ID == 0 {
		return entity.Product{}, nil
	}

	Deleted, err := s.productRepository.Delete(ID)

	if err != nil {
		return entity.Product{}, err
	}

	return Deleted, nil
}
