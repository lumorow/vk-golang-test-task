package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/sirupsen/logrus"
)

// CreateActor creates a new actor in the system.
// @Summary Create actor
// @Description Creates a new actor.
// @Tags Actors
// @Accept  json
// @Produce  json
// @Param actor body entity.Actor true "Data of the new actor"
// @Security ApiKeyAuth
// @Success 200 {integer} integer "ID of the created actor"
// @Failure 400 {string} string "Invalid request data"
// @Failure 500 {string} string "Internal server error"
// @Router /api/actor [post]
func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	err = checkAdminRule(w, r)
	if err != nil {
		return
	}

	var input entity.Actor

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	input.Name = slug.Make(input.Name)

	id, err := h.services.CreateActor(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d add actor with id: %d", userId, id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}
