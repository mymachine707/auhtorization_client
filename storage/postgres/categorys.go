package postgres

import (
	"errors"
	"mymachine707/protogen/eCommerce"
	"time"
)

var err error

// AddCategory ...
func (stg Postgres) AddCategory(id string, entity *eCommerce.CreateCategoryRequest) error {

	_, err = stg.db.Exec(`INSERT INTO category (
		"id",
		"category_name",
		"description"
		) VALUES(
		$1,
		$2,
		$3
	)`,
		id,
		entity.CategoryName,
		entity.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetCategoryByID ...
func (stg Postgres) GetCategoryByID(id string) (*eCommerce.Category, error) {
	result := &eCommerce.Category{}

	var updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT
		"id",
		"category_name",
		"description",
		"created_at",
		"updated_at"
	FROM "category" WHERE "deleted_at" is null AND id=$1`, id).Scan(
		&result.Id,
		&result.CategoryName,
		&result.Description,
		&result.CreatedAt, // !
		&updatedAt,
	)

	if err != nil {
		return result, err
	}

	if updatedAt != nil {
		result.UpdatedAt = updatedAt.String()
	}

	return result, nil
}

// GetCategoryList ...
func (stg Postgres) GetCategoryList(offset, limit int, search string) (resp *eCommerce.GetCategoryListResponse, err error) {

	resp = &eCommerce.GetCategoryListResponse{
		Categorys: make([]*eCommerce.Category, 0),
	}

	rows, err := stg.db.Queryx(`
	Select 
	"id",
	"category_name",
	"description",
	"created_at",
	"updated_at"
 from category WHERE deleted_at is null AND (("category_name" ILIKE '%' || $1 || '%') OR 
		("description" ILIKE '%' || $1 || '%'))
		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &eCommerce.Category{}
		var updatedAt *string

		err = rows.Scan(
			&a.Id,
			&a.CategoryName,
			&a.Description,
			&a.CreatedAt,
			&updatedAt,
		)

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		if err != nil {
			return resp, err
		}

		resp.Categorys = append(resp.Categorys, a)

	}

	return resp, nil
}

// UpdateCategory ...
func (stg Postgres) UpdateCategory(category *eCommerce.UpdateCategoryRequest) error {

	rows, err := stg.db.NamedExec(`Update category set "category_name"=:c, "description"=:d, "updated_at"=now() Where "id"=:id and "deleted_at" is null`, map[string]interface{}{
		"id": category.Id,
		"c":  category.CategoryName,
		"d":  category.Description,
	})

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("category not found")
}

// DeleteCategory ...
func (stg Postgres) DeleteCategory(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE category SET "deleted_at"=now() Where id=$1 and "deleted_at" is null`, idStr)

	if err != nil {
		return err
	}

	n, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("Cannot delete Category becouse Category not found")
}
