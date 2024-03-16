package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

// GetFilmsWithSort handles the request to retrieve films sorted by the specified criteria and associated with the provided actor IDs.
// @Summary Retrieve films sorted by criteria and associated with actors
// @Description This endpoint retrieves films sorted by the specified criteria and associated with the provided actor IDs.
// @Tags Films
// @Param sortType query string false "Sort type: rating, date, name."
// @Param id query []int true "Actor IDs"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} []entity.Film "OK"
// @Failure 400 {string} string "Invalid request or invalid actor ID"
// @Failure 500 {string} string "Internal server error"
// @Router /api/films/sorted [get]
func (h *Handler) GetFilmsWithSort(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	sortType := r.URL.Query().Get("sortType")

	filmsIdStr := r.URL.Query().Get("id")

	if filmsIdStr == "" {
		newErrorResponse(w, http.StatusBadRequest, "missing 'id' parameter")
		return
	}

	filmsIdStrSlice := strings.Split(filmsIdStr, ",")
	filmsIds := make([]int, 0, len(filmsIdStrSlice))

	for _, id := range filmsIdStrSlice {
		actorId, err := strconv.Atoi(id)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid actors id param: %v", err.Error()))
			return
		}
		filmsIds = append(filmsIds, actorId)
	}

	res, err := h.services.GetFilmsWithSort(sortType, filmsIds)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d get films based on sort: %s", userId, sortType)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
