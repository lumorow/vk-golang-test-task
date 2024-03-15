package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

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
	Username string `json:"username"`
	Password string `json:"password"`
}

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
