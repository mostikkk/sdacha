package handlers

import (
	"context"
	"obuch/internal/userService"
	"obuch/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// GetUser реализует users.StrictServerInterface.
func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Pass,
		}
		response = append(response, user)
	}

	return response, nil
}

// PostUser реализует users.StrictServerInterface.
func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.Users{
		Email: *userRequest.Email,
		Pass:  *userRequest.Password,
	}

	createdUser, err := h.Service.PostUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Pass,
	}

	return response, nil
}

// DeleteUser реализует users.StrictServerInterface.
func (h *UserHandler) DeleteUsers(ctx context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	id := request.Params.UserId

	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}

	response := users.DeleteUsers204Response{}

	return response, nil
}

// PutUser реализует users.StrictServerInterface.
func (h *UserHandler) PutUsers(_ context.Context, request users.PutUsersRequestObject) (users.PutUsersResponseObject, error) {
	id := *request.Body.Id
	userRequest := request.Body

	userToUpdate := userService.Users{
		Email: *userRequest.Email,
		Pass:  *userRequest.Password,
	}

	updatedUser, err := h.Service.PatchUserByID(uint(id), userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PutUsers200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Pass,
	}

	return response, nil
}
