package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// GetFilmsWithFragment handles the request to search for films based on actor name and film name fragments.
// @Summary Search for films by actor name and film name fragments
// @Description This endpoint allows searching for films based on actor name and film name fragments.
// @Tags Films
// @Param actorNameFr query string true "Actor name fragment"
// @Param filmNameFr query string true "Film name fragment"
// @Accept json
// @Produce json
// @Success 200 {array} Film "OK"
// @Failure 400 {string} string "Invalid request or invalid actor ID"
// @Failure 500 {string} string "Internal server error"
// @Router /films [get]
func (h *Handler) GetFilmsWithFragment(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(w, r)
	if err != nil {
		return
	}

	actorNameFrag := r.URL.Query().Get("actorNameFr")
	filmNameFrag := r.URL.Query().Get("filmNameFr")

	res, err := h.services.GetFilmWithFragment(actorNameFrag, filmNameFrag)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("user id: %d get films based on actor name and film name: %s, fragments: %s", userId, actorNameFrag, filmNameFrag)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
