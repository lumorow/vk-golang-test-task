package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/sirupsen/logrus"
)

// CreateFilm creates a new film in the system.
// @Summary Create film
// @Description Creates a new film.
// @Tags Films
// @Accept  json
// @Produce  json
// @Param film body entity.Film true "Data of the new film"
// @Security ApiKeyAuth
// @Success 200 {integer} integer "ID of the created film"
// @Failure 400 {string} string "Invalid request data"
// @Failure 500 {string} string "Internal server error"
// @Router /api/film [post]
func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	err = checkAdminRule(w, r)
	if err != nil {
		return
	}

	var input entity.Film

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	input.Name = slug.Make(input.Name)

	id, err := h.services.CreateFilm(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d add film with id: %d", userId, id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}
