package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"net/http"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/sirupsen/logrus"
)

// UpdateActorById updates an actor by its ID in the system.
// @Summary Update actor by ID
// @Description Updates an actor with the specified ID based on the data passed in the request body.
// @Tags Actors
// @Param id path integer true "Actor ID to update"
// @Accept  json
// @Produce  json
// @Param actor body entity.UpdateActorInput true "Data of the actor to update"
// @Security ApiKeyAuth
// @Success 200 {string} string "Actor updated successfully"
// @Failure 400 {string} string "Invalid actor ID param or request data"
// @Failure 500 {string} string "Internal server error"
// @Router /api/actor/{id} [patch]
func (h *Handler) UpdateActorById(w http.ResponseWriter, r *http.Request) {
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

	var input entity.UpdateActorInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	*input.Name = slug.Make(*input.Name)

	if err := h.Service.UpdateActorById(actorId, input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d update actor with id: %d", userId, actorId)
	w.WriteHeader(http.StatusOK)
}
