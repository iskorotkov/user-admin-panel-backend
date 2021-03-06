/*
 * API for user admin panel
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: korotkov.ivan.s@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/iskorotkov/user-admin-panel-backend/entities"
)

// UsersApiService is a service that implements the logic for the UsersApiServicer
// This service should implement the business logic for every endpoint for the UsersApi API.
// Include any external packages or services that will be required by this service.
type UsersApiService struct {
	db *gorm.DB
}

// NewUsersApiService creates a default api service
func NewUsersApiService(db *gorm.DB) UsersApiServicer {
	return &UsersApiService{db: db}
}

// All -
func (s *UsersApiService) All(ctx context.Context) (ImplResponse, error) {
	var users []entities.User
	if err := s.db.Find(&users).Order("id asc").Error; err != nil {
		return Response(
			http.StatusInternalServerError, Error{
				Code:    1001,
				Message: "Error while getting users",
				Errors:  []string{err.Error()},
			},
		), nil
	}

	res := make([]User, 0, len(users))
	for _, user := range users {
		res = append(res, User{
			Id: int32(user.ID),
			NewUser: NewUser{
				Name:   user.Name,
				Phone:  user.Phone,
				Email:  user.Email,
				Gender: string(user.Gender),
			},
		})
	}

	return Response(http.StatusOK, res), nil
}

// Create -
func (s *UsersApiService) Create(ctx context.Context, newUser NewUser) (ImplResponse, error) {
	userToCreate := entities.User{
		Name:   newUser.Name,
		Phone:  newUser.Phone,
		Email:  newUser.Email,
		Gender: entities.Gender(newUser.Gender),
	}

	userToCreate = userToCreate.Trim()

	if errs := userToCreate.Validate(); errs != nil {
		return Response(
			http.StatusUnprocessableEntity, Error{
				Code:    2001,
				Message: "User has validation errors",
				Errors:  errs.Slice(),
			},
		), nil
	}

	if err := s.db.Create(&userToCreate).Error; err != nil {
		return Response(
			http.StatusUnprocessableEntity, Error{
				Code:    2002,
				Message: "Error while creating user",
				Errors:  []string{err.Error()},
			},
		), nil
	}

	return Response(
		http.StatusCreated, User{
			Id:      int32(userToCreate.ID),
			NewUser: newUser,
		},
	), nil
}

// Delete -
func (s *UsersApiService) Delete(ctx context.Context, id int32) (ImplResponse, error) {
	var user entities.User
	if err := s.db.Find(&user, id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return Response(http.StatusInternalServerError, nil), err
	}

	if err := s.db.Where("id = ?", id).Delete(&entities.User{}).Error; err != nil {
		return Response(
			http.StatusInternalServerError, Error{
				Code:    3001,
				Message: "Error while deleting user",
				Errors:  []string{err.Error()},
			},
		), nil
	}

	res := User{
		Id: int32(id),
		NewUser: NewUser{
			Name:   user.Name,
			Phone:  user.Phone,
			Email:  user.Email,
			Gender: string(user.Gender),
		},
	}

	return Response(http.StatusOK, res), nil
}

// Single -
func (s *UsersApiService) Single(ctx context.Context, id int32) (ImplResponse, error) {
	var user entities.User
	if err := s.db.Find(&user, id).Error; err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	if user.ID == 0 {
		return Response(
			http.StatusNotFound, Error{
				Code:    4001,
				Message: "User not found",
				Errors:  []string{"record not found"},
			},
		), nil
	}

	res := User{
		Id: int32(user.ID),
		NewUser: NewUser{
			Name:   user.Name,
			Phone:  user.Phone,
			Email:  user.Email,
			Gender: string(user.Gender),
		},
	}

	return Response(http.StatusOK, res), nil
}

// Update -
func (s *UsersApiService) Update(ctx context.Context, id int32, user User) (ImplResponse, error) {
	if id != user.Id {
		return Response(http.StatusUnprocessableEntity, Error{
			Code:    5001,
			Message: "User id mismatch",
			Errors:  []string{"id mismatch"},
		}), nil
	}

	var userToUpdate entities.User
	if err := s.db.Find(&userToUpdate, id).Error; err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	if userToUpdate.ID == 0 {
		userToUpdate.ID = uint(user.Id)
	}

	userToUpdate = entities.User{
		Model:  userToUpdate.Model,
		Name:   user.Name,
		Phone:  user.Phone,
		Email:  user.Email,
		Gender: entities.Gender(user.Gender),
	}

	userToUpdate = userToUpdate.Trim()

	if errs := userToUpdate.Validate(); errs != nil {
		return Response(
			http.StatusUnprocessableEntity, Error{
				Code:    5002,
				Message: "User has validation errors",
				Errors:  errs.Slice(),
			},
		), nil
	}

	if err := s.db.Save(&userToUpdate).Error; err != nil {
		return Response(
			http.StatusUnprocessableEntity, Error{
				Code:    5003,
				Message: "Error while creating user",
				Errors:  []string{err.Error()},
			},
		), nil
	}

	// User was updated.
	if user.Id != 0 {
		return Response(http.StatusOK, user), nil
	}

	// User was created.
	return Response(
		http.StatusCreated, User{
			Id:      int32(userToUpdate.ID),
			NewUser: user.NewUser,
		},
	), nil
}
