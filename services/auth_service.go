package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mymachine707/protogen/eCommerce"
	"mymachine707/util"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *clientService) Login(ctx context.Context, req *eCommerce.LoginAuthRequest) (*eCommerce.TokenResponse, error) {
	log.Println("Login...")
	fmt.Println(req)
	errUser := errors.New("Username or Password wrong")
	// requsetdan kevotgan username bilan bazadigi username qidirib topiladi
	user, err := s.stg.GetClientByUsername(req.Username)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, errUser.Error())
	}

	// requsetdan kevotgan password bilan bazadan kegan username passwordini solishtiriladi.
	match, err := util.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, " s.stg.GetUserByUsername: %s", err.Error())
	}

	if !match {
		return nil, status.Errorf(codes.Unauthenticated, errUser.Error())
	}

	//Token generator
	m := map[string]interface{}{
		"user_id":  user.Id,
		"username": user.Username,
	}

	token, err := util.GenerateJWT(m, 1*time.Hour, s.cfg.SECRET_KEY)
	if err != nil {
		return nil, status.Errorf(codes.Internal, " util.GenerateJWT: %s", err.Error())
	}

	return &eCommerce.TokenResponse{
		Token: token,
	}, nil
}

func (s *clientService) HasAccess(ctx context.Context, req *eCommerce.TokenRequest) (*eCommerce.HasAccessResponse, error) {
	log.Println("HasAccess...")

	result, err := util.ParseClaims(req.Token, s.cfg.SECRET_KEY)
	if err != nil {
		log.Println(status.Errorf(codes.PermissionDenied, " util.ParseClaims: %s", err.Error()))
		return &eCommerce.HasAccessResponse{
			User:      nil,
			HasAccess: false,
		}, nil
	}

	log.Println(result.Username)

	user, err := s.stg.GetClientByID(result.UserID)
	if err != nil {
		log.Println(status.Errorf(codes.PermissionDenied, " util.ParseClaims: %s", err.Error()))
		return &eCommerce.HasAccessResponse{
			User:      user,
			HasAccess: false,
		}, nil
	}

	return &eCommerce.HasAccessResponse{
		User:      user,
		HasAccess: true,
	}, nil
}
