package handler

import (
	"net/http"
	"strconv"
)

// DeleteFilmById deletes a film by its ID from the system.
// @Summary Delete film by ID
// @Description Deletes a film with the specified ID.
// @Tags Films
// @Param id path integer true "Film ID to delete"
// @Success 200 {string} string "Film deleted successfully"
// @Failure 400 {string} string "Invalid film ID param"
// @Failure 500 {string} string "Internal server error"
// @Router /film/{id} [delete]
func (h *Handler) DeleteFilmById(w http.ResponseWriter, r *http.Request) {
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

	if err := h.services.DeleteFilmById(filmId); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
