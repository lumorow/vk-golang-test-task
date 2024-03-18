package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// GetActorsWithFilms returns actors with their associated films based on the provided actor IDs.
// @Summary Get actors with associated films
// @Description Returns actors with their associated films based on the provided actor IDs.
// @Tags Actors
// @Param id query []int true "Actor IDs"
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} []entity.ActorFilms "OK"
// @Failure 400 {string} string "Invalid request or invalid actor ID"
// @Failure 500 {string} string "Internal server error"
// @Router /api/actors [get]
func (h *Handler) GetActorsWithFilms(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	actorsIdStr := r.URL.Query().Get("id")
	if actorsIdStr == "" {
		newErrorResponse(w, http.StatusBadRequest, "missing 'id' parameter")
		return
	}

	actorsIdStrSlice := strings.Split(actorsIdStr, ",")
	actorsId := make([]int, len(actorsIdStrSlice))

	for i, idStr := range actorsIdStrSlice {
		actorId, err := strconv.Atoi(idStr)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid actors id param: %v", err))
			return
		}
		actorsId[i] = actorId
	}

	res, err := h.ActorService.GetActorsWithFilms(actorsId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d get actors: %d with associated films", userId, actorsId)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
