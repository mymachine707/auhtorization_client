package products

import (
	"context"
	"fmt"
	"log"
	"mymachine707/protogen/eCommerce"
	"mymachine707/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productService struct {
	stg storage.Interfaces
	eCommerce.UnimplementedProductServiceServer
}

func NewProductService(stg storage.Interfaces) *productService {
	return &productService{
		stg: stg,
	}
}

func (s *productService) Ping(ctx context.Context, req *eCommerce.Empty) (*eCommerce.Pong, error) {
	log.Println("Ping")

	return &eCommerce.Pong{
		Message: "Ok",
	}, nil
}

func (s *productService) CreateProduct(ctx context.Context, req *eCommerce.CreateProductRequest) (*eCommerce.Product, error) {
	fmt.Println("<<< ---- CreateProduct ---->>>")
	// create new product
	id := uuid.New()

	err := s.stg.AddProduct(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddProduct: %s", err.Error())
	}

	product, err := s.stg.GetProductByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductByID: %s", err.Error())
	}

	return &eCommerce.Product{
		Id:          product.Id,
		CategoryId:  product.Category.Id,
		ProductName: product.ProductName,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *productService) UpdateProduct(ctx context.Context, req *eCommerce.UpdateProductRequest) (*eCommerce.Product, error) {
	fmt.Println("<<< ---- UpdateProduct ---->>>")

	err := s.stg.UpdateProduct(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateProduct: %s", err.Error())
	}

	product, err := s.stg.GetProductByID(req.Id) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductByID---!: %s", err.Error())
	}

	return &eCommerce.Product{
		Id:          product.Id,
		CategoryId:  product.Category.Id,
		ProductName: product.ProductName,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *productService) DeleteProduct(ctx context.Context, req *eCommerce.DeleteProductRequest) (*eCommerce.Product, error) {
	fmt.Println("<<< ---- DeleteProduct ---- >>>")

	product, err := s.stg.GetProductByID(req.Id) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductByID: %s", err.Error())
	}

	err = s.stg.DeleteProduct(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteProduct: %s", err.Error())
	}

	return &eCommerce.Product{
		Id:          product.Id,
		CategoryId:  product.Category.Id,
		ProductName: product.ProductName,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *productService) GetProductList(ctx context.Context, req *eCommerce.GetProductListRequest) (*eCommerce.GetProductListResponse, error) {
	fmt.Println("<<< ---- GetProductList ---->>>")

	res, err := s.stg.GetProductList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductList: %s", err.Error())
	}

	return res, nil
}

func (s *productService) GetProductById(ctx context.Context, req *eCommerce.GetProductByIDRequest) (*eCommerce.GetProductByIDResponse, error) {
	fmt.Println("<<< ---- GetProductById ---->>>")

	product, err := s.stg.GetProductByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetProductByID: %s", err.Error())
	}

	return product, nil
}
