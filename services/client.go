package client

import (
	"context"
	"fmt"
	"log"

	"mymachine707/config"
	"mymachine707/protogen/eCommerce"
	"mymachine707/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type clientService struct {
	cfg config.Config
	stg storage.Interfaces
	eCommerce.UnimplementedClientServiceServer
}

func NewClientService(stg storage.Interfaces) *clientService {
	return &clientService{
		stg: stg,
	}
}
func (s *clientService) Ping(ctx context.Context, req *eCommerce.Empty) (*eCommerce.Pong, error) {
	fmt.Println("<<< ---- Ping ---->>>")
	log.Println("Ping")
	return &eCommerce.Pong{
		Message: "Ok",
	}, nil
}

func (s *clientService) CreateClient(ctx context.Context, req *eCommerce.CreateClientRequest) (*eCommerce.Client, error) {
	fmt.Println("<<< ---- CreateClient ---->>>")

	id := uuid.New()

	err := s.stg.AddClient(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddClient: %s", err)
	}

	client, err := s.stg.GetClientByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetClientByID: %s", err)
	}

	return &eCommerce.Client{
		Id:          client.Id,
		Firstname:   client.Firstname,
		Lastname:    client.Lastname,
		Username:    client.Username,
		PhoneNumber: client.PhoneNumber,
		Address:     client.Address,
		Type:        client.Type,
		Password:    client.Password,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
	}, nil
}

func (s *clientService) UpdateClient(ctx context.Context, req *eCommerce.UpdateClientRequest) (*eCommerce.Client, error) {
	fmt.Println("<<< ---- UpdateClient ---->>>")

	err := s.stg.UpdateClient(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateClient: %s", err)
	}

	client, err := s.stg.GetClientByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetClientByID: %s", err)
	}

	return &eCommerce.Client{
		Id:          client.Id,
		Firstname:   client.Firstname,
		Lastname:    client.Lastname,
		Username:    client.Username,
		PhoneNumber: client.PhoneNumber,
		Address:     client.Address,
		Type:        client.Type,
		Password:    client.Password,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
	}, nil
}

func (s *clientService) DeleteClient(ctx context.Context, req *eCommerce.DeleteClientRequest) (*eCommerce.Client, error) {
	fmt.Println("<<< ---- DeleteClient ---->>>")

	client, err := s.stg.GetClientByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetClientByID: %s", err)
	}

	err = s.stg.DeleteClient(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteClient: %s", err)
	}

	return &eCommerce.Client{
		Id:          client.Id,
		Firstname:   client.Firstname,
		Lastname:    client.Lastname,
		Username:    client.Username,
		PhoneNumber: client.PhoneNumber,
		Address:     client.Address,
		Type:        client.Type,
		Password:    client.Password,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
	}, nil
}

func (s *clientService) GetClientList(ctx context.Context, req *eCommerce.GetClientListRequest) (*eCommerce.GetClientListResponse, error) {
	fmt.Println("<<< ---- GetClientList ---->>>")

	res, err := s.stg.GetClientList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetClientList: %s", err)
	}

	return res, nil
}

func (s *clientService) GetClientById(ctx context.Context, req *eCommerce.GetClientByIDRequest) (*eCommerce.Client, error) {
	fmt.Println("<<< ---- GetClientById ---->>>")

	client, err := s.stg.GetClientByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetClientByID: %s", err)
	}

	return &eCommerce.Client{
		Id:          client.Id,
		Firstname:   client.Firstname,
		Lastname:    client.Lastname,
		Username:    client.Username,
		PhoneNumber: client.PhoneNumber,
		Address:     client.Address,
		Type:        client.Type,
		Password:    client.Password,
		CreatedAt:   client.CreatedAt,
		UpdatedAt:   client.UpdatedAt,
	}, nil
}
