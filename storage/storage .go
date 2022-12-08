package storage

import (
	"mymachine707/protogen/eCommerce"
)

// Interfaces ...
type Interfaces interface {
	// For Product
	AddProduct(id string, entity *eCommerce.CreateProductRequest) error
	GetProductByID(id string) (*eCommerce.GetProductByIDResponse, error)
	GetProductList(offset, limit int, search string) (*eCommerce.GetProductListResponse, error)
	UpdateProduct(product *eCommerce.UpdateProductRequest) error
	DeleteProduct(idStr string) error

	// For Category
	AddCategory(id string, entity *eCommerce.CreateCategoryRequest) error
	GetCategoryByID(id string) (*eCommerce.Category, error)
	GetCategoryList(offset, limit int, search string) (resp *eCommerce.GetCategoryListResponse, err error)
	UpdateCategory(category *eCommerce.UpdateCategoryRequest) error
	DeleteCategory(idStr string) error
}
