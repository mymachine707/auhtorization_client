package category

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

type categoryService struct {
	stg storage.Interfaces
	eCommerce.UnimplementedCategoryServiceServer
}

func NewCategoryService(stg storage.Interfaces) *categoryService {
	return &categoryService{
		stg: stg,
	}
}
func (s *categoryService) Ping(ctx context.Context, req *eCommerce.Empty) (*eCommerce.Pong, error) {
	fmt.Println("<<< ---- Ping ---->>>")
	log.Println("Ping")
	return &eCommerce.Pong{
		Message: "Ok",
	}, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, req *eCommerce.CreateCategoryRequest) (*eCommerce.Category, error) {
	fmt.Println("<<< ---- CreateCategory ---->>>")

	id := uuid.New()

	err := s.stg.AddCategory(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddCategory: %s", err)
	}

	category, err := s.stg.GetCategoryByID(id.String()) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	return &eCommerce.Category{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		Description:  category.Description,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, req *eCommerce.UpdateCategoryRequest) (*eCommerce.Category, error) {
	fmt.Println("<<< ---- UpdateCategory ---->>>")

	err := s.stg.UpdateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateCategory: %s", err)
	}

	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	return &eCommerce.Category{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		Description:  category.Description,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, req *eCommerce.DeleteCategoryRequest) (*eCommerce.Category, error) {
	fmt.Println("<<< ---- DeleteCategory ---->>>")

	category, err := s.stg.GetCategoryByID(req.Id) // maqsad tekshirish rostan  ham create bo'ldimi?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	err = s.stg.DeleteCategory(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteCategory: %s", err)
	}

	return &eCommerce.Category{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		Description:  category.Description,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}, nil
}

func (s *categoryService) GetCategoryList(ctx context.Context, req *eCommerce.GetCategoryListRequest) (*eCommerce.GetCategoryListResponse, error) {
	fmt.Println("<<< ---- GetCategoryList ---->>>")

	res, err := s.stg.GetCategoryList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryList: %s", err)
	}

	return res, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, req *eCommerce.GetCategoryByIDRequest) (*eCommerce.Category, error) {
	fmt.Println("<<< ---- GetCategoryById ---->>>")

	category, err := s.stg.GetCategoryByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetCategoryByID: %s", err)
	}

	return &eCommerce.Category{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		Description:  category.Description,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}, nil
}
