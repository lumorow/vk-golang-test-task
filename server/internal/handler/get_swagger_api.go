package handler

import (
	_ "filmlib/server/docs"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetSwaggerAPI(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("server/docs/swagger.yml")).ServeHTTP(w, r)
	logrus.Print("open Swagger UI")
	return
}
