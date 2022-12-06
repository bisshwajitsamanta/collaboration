package redbus

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingUser   = errors.New("Failed to Fetch the user")
	ErrNotImplemented = errors.New("Not Implemented Yet")
)

// Store - This interface defines all methods that our applications need
type Store interface {
	GetUser(context.Context, string) (UserLogin, error)
}

//UserLogin - A representation of Login to Redbus structure of our service
type UserLogin struct {
	ID       string
	UserName string
}

//Service - All our Logic will be built on top of it
type Service struct {
	Store Store
}

//NewService - Accepting interface and returning struct
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

//GetUser - Retrieves user details by checking out from DB
func (s *Service) GetUser(ctx context.Context, id string) (UserLogin, error) {
	fmt.Println("Retrieving user details")
	cmt, err := s.Store.GetUser(ctx, id)
	if err != nil {
		fmt.Println(err)
		return UserLogin{}, ErrFetchingUser
	}
	return cmt, nil
}

// NewUser - Creates new user
func (s *Service) NewUser(ctx context.Context, user UserLogin) (UserLogin, error) {
	return UserLogin{}, ErrNotImplemented
}
