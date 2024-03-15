package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"github.com/gosimple/slug"
	"net/http"
)

// CreateFilm creates a new film in the system.
// @Summary Create film
// @Description Creates a new film.
// @Tags Films
// @Accept  json
// @Produce  json
// @Param film body FilmInput true "Data of the new film"
// @Success 200 {integer} integer "ID of the created film"
// @Failure 400 {string} string "Invalid request data"
// @Failure 500 {string} string "Internal server error"
// @Router /film [post]
func (h *Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}
