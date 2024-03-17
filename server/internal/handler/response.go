package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
