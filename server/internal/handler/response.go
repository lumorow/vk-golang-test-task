package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
