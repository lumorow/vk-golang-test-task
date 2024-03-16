package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// DeleteActorById deletes an actor by its ID from the system.
// @Summary Delete actor by ID
// @Description Deletes an actor with the specified ID.
// @Tags Actors
// @Param id path integer true "Actor ID to delete"
// @Security ApiKeyAuth
// @Success 200 {string} string "Actor deleted successfully"
// @Failure 400 {string} string "Invalid actor ID param"
// @Failure 500 {string} string "Internal server error"
// @Router /actor/{id} [delete]
func (h *Handler) DeleteActorById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	err = checkAdminRule(w, r)
	if err != nil {
		return
	}

	matches := ActorReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		newErrorResponse(w, http.StatusBadRequest, "invalid actor id param")
		return
	}

	actorId, err := strconv.Atoi(matches[1])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid actor id param")
		return
	}

	if err := h.services.DeleteActorById(actorId); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d delete actor with id: %d", userId, actorId)
	w.WriteHeader(http.StatusOK)
}
