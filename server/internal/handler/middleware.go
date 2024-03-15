package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func userIdentity(h *Handler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get(authorizationHeader)
			if header == "" {
				newErrorResponse(w, http.StatusUnauthorized, "пустой заголовок авторизации")
				return
			}

			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 {
				newErrorResponse(w, http.StatusUnauthorized, "неверный заголовок авторизации")
				return
			}

			userId, err := h.services.ParseToken(headerParts[1])
			if err != nil {
				newErrorResponse(w, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), userCtx, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//func getUserId(w http.ResponseWriter, r *http.Request) (int, error) {
//	id, ok := c.Get(userCtx)
//	if !ok {
//		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
//		return 0, errors.New("user id not found")
//	}
//
//	idInt, ok := id.(int)
//	if !ok {
//		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
//		return 0, errors.New("user id not found")
//	}
//
//	return idInt, nil
//}
