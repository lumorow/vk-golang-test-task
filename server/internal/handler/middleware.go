package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userRoleHeader      = "userRole"
	UserIdHeader        = "userId"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, userRole, err := h.services.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), UserIdHeader, userId)
		ctx = context.WithValue(ctx, userRoleHeader, userRole)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserId(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.Header.Get(UserIdHeader)
	if len(id) == 0 {
		newErrorResponse(w, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func checkAdminRule(w http.ResponseWriter, r *http.Request) error {
	role := r.Header.Get(userRoleHeader)

	if role != "admin" {
		newErrorResponse(w, http.StatusInternalServerError, "not enough rights")
		return errors.New("not enough rights")
	}

	return nil
}
