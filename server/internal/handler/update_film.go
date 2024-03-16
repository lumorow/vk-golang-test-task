package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"github.com/gosimple/slug"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// UpdateFilmById updates a film by its ID in the system.
// @Summary Update film by ID
// @Description Updates a film with the specified ID based on the data passed in the request body.
// @Tags Films
// @Param id path integer true "Film ID to update"
// @Accept  json
// @Produce  json
// @Param film body entity.UpdateFilmInput true "Data of the film to update"
// @Security ApiKeyAuth
// @Success 200 {string} string "Film updated successfully"
// @Failure 400 {string} string "Invalid film ID param or request data"
// @Failure 500 {string} string "Internal server error"
// @Router /api/film/{id} [patch]
func (h *Handler) UpdateFilmById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	err = checkAdminRule(w, r)
	if err != nil {
		return
	}

	matches := FilmReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		newErrorResponse(w, http.StatusBadRequest, "invalid film id param")
		return
	}

	filmId, err := strconv.Atoi(matches[1])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid film id param")
		return
	}

	var input entity.UpdateFilmInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	*input.Name = slug.Make(*input.Name)

	if err := h.services.UpdateFilmById(filmId, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d update film with id: %d", userId, filmId)
	w.WriteHeader(http.StatusOK)
}
