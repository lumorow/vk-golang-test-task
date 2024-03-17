package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// @Summary Sign up a new user
// @Description Creates a new user account.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body entity.User true "User data example: '{'username': 'example_user', 'password': 'example_password', 'role': 'Enums(admin, user)'}"
// @Success 200 {string} string "ID of the created user"
// @Failure 400 {string} string "Invalid request data"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input entity.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("sign up user with id: %d", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}

type signInInput struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
}

// @Summary Sign in a user
// @Description Signs in an existing user.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body signInInput true "User credentials"
// @Success 200 {string} string "JWT token"
// @Failure 400 {string} string "Invalid request data"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("sign in user with name: %s", input.Username)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("token: %s", token)))
}
