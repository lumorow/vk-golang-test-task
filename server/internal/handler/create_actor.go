package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"github.com/gosimple/slug"
	"net/http"
)

// CreateActor creates a new actor in the system.
// @Summary Create actor
// @Description Creates a new actor.
// @Tags Actors
// @Accept  json
// @Produce  json
// @Param actor body ActorInput true "Data of the new actor"
// @Success 200 {integer} integer "ID of the created actor"
// @Failure 400 {string} string "Invalid request data"
// @Failure 500 {string} string "Internal server error"
// @Router /actor [post]
func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}
