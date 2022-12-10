package storage

import (
	"mymachine707/protogen/eCommerce"
)

// Interfaces ...
type Interfaces interface {
	// For Client
	AddClient(id string, entity *eCommerce.CreateClientRequest) error
	GetClientByID(id string) (*eCommerce.Client, error)
	GetClientList(offset, limit int, search string) (resp *eCommerce.GetClientListResponse, err error)
	UpdateClient(client *eCommerce.UpdateClientRequest) error
	DeleteClient(idStr string) error

	GetClientByUsername(username string) (*eCommerce.Client, error)
}
