package service

import (
	"context"

	categoryPb "github.com/jihadable/learn-grpc-proto/category"
	productPb "github.com/jihadable/learn-grpc-proto/product"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
}

func (service *ProductService) CreateProduct(ctx context.Context, req *productPb.Product) (*productPb.Product, error) {
	product := &productPb.Product{
		Id:          2,
		Name:        "Product 1",
		Description: "Desc 4",
		Price:       34000,
		Category: &categoryPb.Category{
			Id:   3,
			Name: "Category 6",
		},
	}

	return product, nil
}

func (service *ProductService) GetProducts(ctx context.Context, _ *productPb.Empty) (*productPb.Products, error) {
	products := &productPb.Products{}
	product := &productPb.Product{
		Id:          6,
		Name:        "Product 4",
		Description: "Desc 9",
		Price:       55000,
		Category: &categoryPb.Category{
			Id:   9,
			Name: "Category 4",
		},
	}

	products.Data = append(products.Data, product)

	return products, nil
}

func (service *ProductService) GetProduct(ctx context.Context, id *productPb.Id) (*productPb.Product, error) {
	product := &productPb.Product{
		Id:          7,
		Name:        "Product 2",
		Description: "Desc 5",
		Price:       16000,
		Category: &categoryPb.Category{
			Id:   2,
			Name: "Category 1",
		},
	}

	return product, nil
}

func (service *ProductService) UpdateProduct(ctx context.Context, id *productPb.Id) (*productPb.Product, error) {
	product := &productPb.Product{
		Id:          3,
		Name:        "Product 4",
		Description: "Desc 2",
		Price:       111000,
		Category: &categoryPb.Category{
			Id:   8,
			Name: "Category 9",
		},
	}

	return product, nil
}

func (service *ProductService) DeleteProduct(ctx context.Context, id *productPb.Id) (*productPb.Status, error) {
	status := &productPb.Status{
		Status: 1,
	}

	return status, nil
}

func NewProductService() *ProductService {
	return &ProductService{}
}
