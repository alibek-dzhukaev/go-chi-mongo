package usecase

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/alibek-dzhukaev/go-abs-beg/dto"
	"github.com/alibek-dzhukaev/go-abs-beg/model"
	"github.com/alibek-dzhukaev/go-abs-beg/repository"
	"github.com/go-chi/chi/v5"
)

type UserService struct {
	DBClient repository.UserInterface
}

func (srv UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	// extract body from request
	var userReq dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "invalid request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	user := model.User{
		Name:    userReq.Name,
		Age:     userReq.Age,
		Country: userReq.Country,
	}

	// call db layer

	result, err := srv.DBClient.CreateUser(user)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while inserting user in db"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// success
	slog.Info("user successfully inserted", slog.String("_id", result))
	res.Data = result
	json.NewEncoder(w).Encode(res)
}

func (srv UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("id field is empty")
		res.Error = "invalid id"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := srv.DBClient.GetUserById(id)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "erro while fetching user from db"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// success
	slog.Info("user successfully fetched")
	res.Data = user
	json.NewEncoder(w).Encode(res)
}

func (srv UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	users, err := srv.DBClient.GetAllUsers()
	if err != nil {
		slog.Error(err.Error())
		res.Error = "erro while fetching user from db"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// success
	slog.Info("users successfully fetched")
	res.Data = users
	json.NewEncoder(w).Encode(res)
}

func (srv UserService) UpdateUserAgeById(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("id field is empty")
		res.Error = "invalid id"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// extract body from request
	var userReq dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "invalid request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := srv.DBClient.UpdateUserAgeById(id, userReq.Age)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "error while updating user in db"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// success
	slog.Info("user successfully updated")
	res.Data = result
	json.NewEncoder(w).Encode(res)
}

func (srv UserService) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("id field is empty")
		res.Error = "invalid id"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := srv.DBClient.DeleteUserById(id)
	if err != nil {
		slog.Error(err.Error())
		res.Error = "erro while deleting user in db"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// success
	slog.Info("user successfully deleted")
	res.Data = user
	json.NewEncoder(w).Encode(res)
}

func (srv UserService) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	res := dto.UserResponse{}

	users, err := srv.DBClient.DeleteAllUsers()
	if err != nil {
		slog.Error(err.Error())
		res.Error = "erro while deleting users in db"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// success
	slog.Info("users successfully deleted")
	res.Data = users
	json.NewEncoder(w).Encode(res)
}
