package postgres

import (
	"errors"
	"mymachine707/protogen/eCommerce"
	"time"
)

var err error

// AddClient ...
func (stg Postgres) AddClient(id string, entity *eCommerce.CreateClientRequest) error {

	_, err = stg.db.Exec(`INSERT INTO client (
		"id",
		"firstname",
		"lastname",
		"username",
		"phone",
		"address",
		"type",
		"password"
		) VALUES(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)`,
		id,
		entity.Firstname,
		entity.Lastname,
		entity.Username,
		entity.PhoneNumber,
		entity.Address,
		entity.Type,
		entity.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetClientByID ...
func (stg Postgres) GetClientByID(id string) (*eCommerce.Client, error) {
	result := &eCommerce.Client{}

	var updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT
		"id",
		"firstname",
		"lastname",
		"username",
		"phone",
		"address",
		"type",
		"password",
		"created_at",
		"updated_at"
	FROM client WHERE "deleted_at" is null AND id=$1`, id).Scan(
		&result.Id,
		&result.Firstname,
		&result.Lastname,
		&result.Username,
		&result.PhoneNumber,
		&result.Address,
		&result.Type,
		&result.Password,
		&result.CreatedAt,
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

// GetClientList ...
func (stg Postgres) GetClientList(offset, limit int, search string) (resp *eCommerce.GetClientListResponse, err error) {

	resp = &eCommerce.GetClientListResponse{
		Clients: make([]*eCommerce.Client, 0),
	}

	rows, err := stg.db.Queryx(`
	Select 
	"id",
	"firstname",
	"lastname",
	"username",
	"phone",
	"address",
	"type",
	"password",
	"created_at",
	"updated_at"
 from client WHERE deleted_at is null AND 
 		(
		("firstname" ILIKE '%' || $1 || '%') OR 
		("lastname" ILIKE '%' || $1 || '%') OR 
		("username" ILIKE '%' || $1 || '%') OR 
		("phone" ILIKE '%' || $1 || '%') OR 
		("address" ILIKE '%' || $1 || '%') OR
		("type" ILIKE '%' || $1 || '%'))
		LIMIT $2 
		OFFSET $3`, search, limit, offset)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &eCommerce.Client{}
		var updatedAt *string

		err = rows.Scan(
			&a.Id,
			&a.Firstname,
			&a.Lastname,
			&a.Username,
			&a.PhoneNumber,
			&a.Address,
			&a.Type,
			&a.Password,
			&a.CreatedAt,
			&updatedAt,
		)

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		if err != nil {
			return resp, err
		}

		resp.Clients = append(resp.Clients, a)

	}

	return resp, nil
}

// UpdateClient ...
func (stg Postgres) UpdateClient(client *eCommerce.UpdateClientRequest) error {

	rows, err := stg.db.NamedExec(`Update client set "firstname"=:f, "lastname"=:l,"username"=:u,"phone"=:p,"address"=:a,"type"=:t, "updated_at"=now() Where "id"=:id and "deleted_at" is null`, map[string]interface{}{
		"id": client.Id,
		"f":  client.Firstname,
		"l":  client.Lastname,
		"u":  client.Username,
		"p":  client.PhoneNumber,
		"a":  client.Address,
		"t":  client.Type,
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

	return errors.New("client not found")
}

// DeleteClient ...
func (stg Postgres) DeleteClient(idStr string) error {
	rows, err := stg.db.Exec(`UPDATE client SET "deleted_at"=now() Where id=$1 and "deleted_at" is null`, idStr)

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

	return errors.New("Cannot delete Client becouse Client not found")
}
