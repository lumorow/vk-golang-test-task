package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetActorsWithFilms returns actors with their associated films based on the provided actor IDs.
// @Summary Get actors with associated films
// @Description Returns actors with their associated films based on the provided actor IDs.
// @Tags Actors
// @Param id query []int true "Actor IDs"
// @Accept  json
// @Produce  json
// @Success 200 {array} []entity.ActorFilms "OK"
// @Failure 400 {string} string "Invalid request or invalid actor ID"
// @Failure 500 {string} string "Internal server error"
// @Router /actors [get]
func (h *Handler) GetActorsWithFilms(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	queryValues := r.URL.Query()["id"]

	if len(queryValues) == 0 {
		newErrorResponse(w, http.StatusBadRequest, "missing 'id' parameter")
		return
	}

	actorsId := make([]int, 0, len(queryValues))

	for _, id := range queryValues {
		actorId, err := strconv.Atoi(id)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid actors id param: %v", err.Error()))
			return
		}
		actorsId = append(actorsId, actorId)
	}

	res, err := h.services.GetActorsWithFilms(actorsId)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d get actors: %s with associated films", userId, queryValues)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
