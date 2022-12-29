package domain

import (
	"context"
	"errors"
	"goqlposgress/models"
	"log"
)

// Login is the resolver for the login field.
func (d *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, ErrBadCredentials
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

// Register is the resolver for the register field.
func (d *Domain) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}
	_, err = d.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already in use")
	}
	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	//TODO: create verification code
	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error while creating transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()
	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error while creating transaction: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		log.Printf("error while committing transaction: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}
	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
