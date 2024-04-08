package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	err := Error{
		Message: message,
	}
	logrus.Error(err)

	http.Error(w, err.Message, statusCode)
}
