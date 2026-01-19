package product

import (
	"context"
)

type ProductService struct {
	repo *ProductRepo
}

func NewProductService(repo *ProductRepo) *ProductService {
	return &ProductService{repo}
}

func (service *ProductService) List(c context.Context) ([]Product, error) {
	products, err := service.repo.List(c)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductService) Create(c context.Context, request *ProductRequest) error {
	product := &Product{
		Name:  request.Name,
		Price: request.Price,
		Stock: request.Stock,
	}

	return service.repo.Create(c, product)
}

func (service *ProductService) Show(c context.Context, id int64) (*Product, error) {
	product, err := service.repo.FindById(c, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) Update(c context.Context, id int64, request *ProductRequest) error {
	product, err := service.repo.FindById(c, id)
	if err != nil {
		return err
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Stock = request.Stock

	return service.repo.Update(c, product)
}

func (service *ProductService) Delete(c context.Context, id int64) error {
	product, err := service.repo.FindById(c, id)
	if err != nil {
		return err
	}

	return service.repo.Delete(c, product)
}
