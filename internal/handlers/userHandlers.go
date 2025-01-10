package handlers

import (
	"context"

	userservice "Test.go/internal/userService"
	"Test.go/internal/web/users"
)

type userHandler struct {
	Service *userservice.UserService
}

func (u *userHandler) DeleteUserByID(ctx context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	err := u.Service.DeleteUserByID(request.UserId)
	if err != nil {
		return users.DeleteUserByID204Response{}, nil
	}
	return users.DeleteUserByID204Response{}, nil
}

func (u *userHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return users.GetUsers200JSONResponse{}, err
	}
	var resp users.GetUsers200JSONResponse
	for _, us := range allUsers {
		user := users.User{
			Id:       &us.ID,
			Email:    &us.Email,
			Password: &us.Password,
		}
		resp = append(resp, user)
	}
	return resp, nil
}

func (u *userHandler) PatchUserByID(ctx context.Context, request users.PatchUserByIDRequestObject) (users.PatchUserByIDResponseObject, error) {
	userRequest := request.Body
	userInstance := userservice.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	updatedUser, err := u.Service.UpdateUserByID(request.UserId, userInstance)
	if err != nil {
		return users.PatchUserByID200JSONResponse{}, err
	}
	response := users.PatchUserByID200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil

}

func (u *userHandler) PostUser(ctx context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	requestUser := request.Body
	userToCreate := userservice.User{
		Email:    *requestUser.Email,
		Password: *requestUser.Password,
	}
	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUser201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func NewUserHandler(service *userservice.UserService) *userHandler {
	return &userHandler{Service: service}
}
