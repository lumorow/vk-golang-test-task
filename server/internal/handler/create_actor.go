package handler

import (
	"encoding/json"
	"filmlib/server/internal/entity"
	"fmt"
	"net/http"
)

func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var input entity.Actor

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateActor(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}
