package postgres

import (
	"errors"
	"time"

	// "mymachine707/eCommerce"
	"mymachine707/protogen/eCommerce"
)

// AddProduct ...
func (stg Postgres) AddProduct(id string, entity *eCommerce.CreateProductRequest) error {
	_, err := stg.GetCategoryByID(entity.CategoryId)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO product (
		"id",
		"category_id",
		"product_name",
		"description",
		"price"
		) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
)`,
		id,
		entity.CategoryId,
		entity.ProductName,
		entity.Description,
		entity.Price,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetProductByID ...  //  ????
func (stg Postgres) GetProductByID(id string) (*eCommerce.GetProductByIDResponse, error) {
	// var res eCommerce.GetProductByIDResponse
	res := &eCommerce.GetProductByIDResponse{
		Category: &eCommerce.GetProductByIDResponse_Category{},
	}

	if id == "" {
		return res, errors.New("id must exist")
	}

	var deletedAt *time.Time
	var updatedAt, categoryUpdatedAt *string
	err := stg.db.QueryRow(`SELECT 
    pr."id",
    pr."product_name",
    pr."description",
	pr."price",
    pr."created_at",
    pr."updated_at",
    pr."deleted_at",
    ca."id",
    ca."category_name",
    ca."description",
    ca."created_at",
    ca."updated_at"
FROM "product" AS pr JOIN "category" AS ca ON pr."category_id" = ca."id" WHERE pr."id" = $1`, id).Scan(
		&res.Id,
		&res.ProductName,
		&res.Description,
		&res.Price,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
		&res.Category.Id,
		&res.Category.CategoryName,
		&res.Category.Description,
		&res.Category.CreatedAt,
		&categoryUpdatedAt,
	)

	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if categoryUpdatedAt != nil {
		res.Category.UpdatedAt = *categoryUpdatedAt
	}

	if deletedAt != nil {
		return res, errors.New("product not found")
	}

	return res, err
}

// GetProductList ...
func (stg Postgres) GetProductList(offset, limit int, search string) (*eCommerce.GetProductListResponse, error) {

	resp := &eCommerce.GetProductListResponse{
		Products: make([]*eCommerce.Product, 0),
	}

	rows, err := stg.db.Queryx(`
	SELECT * FROM product WHERE 
	(product_name || ' ' || description ILIKE '%' || $1 || '%') AND "deleted_at" is null
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &eCommerce.Product{}

		var updatedAt, deletedAt *string
		err := rows.Scan(
			&a.Id,
			&a.CategoryId,
			&a.ProductName,
			&a.Description,
			&a.Price,
			&a.CreatedAt,
			&updatedAt,
			&deletedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Products = append(resp.Products, a)
	}

	return resp, err
}

// UpdateProduct ...
func (stg Postgres) UpdateProduct(product *eCommerce.UpdateProductRequest) error {

	res, err := stg.db.NamedExec(`UPDATE product SET "product_name"=:pn, "description"=:d, "price"=:p, "updated_at"=now() WHERE "id"=:id AND "deleted_at" is null`, map[string]interface{}{
		"id": product.Id,
		"pn": product.ProductName,
		"d":  product.Description,
		"p":  product.Price,
	})

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("product not found")
}

// DeleteProduct ...
func (stg Postgres) DeleteProduct(idStr string) error {

	res, err := stg.db.Exec(`UPDATE product Set "deleted_at"=now() WHERE id=$1 AND "deleted_at" is null`, idStr)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("product not found")
}
