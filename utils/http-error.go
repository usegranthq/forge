package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

type HttpErrorUtil struct{}

var HttpError = &HttpErrorUtil{}

func getMessage(code int, message []string) string {
	if len(message) > 0 && message[0] != "" {
		return message[0]
	}
	return http.StatusText(code)
}

func getErrorCode(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "bad_request"
	case http.StatusUnauthorized:
		return "unauthorized"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not_found"
	case http.StatusInternalServerError:
		return "internal_server_error"
	case http.StatusConflict:
		return "conflict"
	default:
		return "unknown_error"
	}
}

func (e *HttpErrorUtil) respondWithError(c *gin.Context, code int, message ...string) {
	errorMessage := getMessage(code, message)
	c.JSON(code, httpError{
		Error:       getErrorCode(code),
		Description: errorMessage,
	})
}

func (e *HttpErrorUtil) InternalServerError(c *gin.Context, message ...string) {
	e.respondWithError(c, http.StatusInternalServerError, message...)
}

func (e *HttpErrorUtil) BadRequest(c *gin.Context, message ...string) {
	e.respondWithError(c, http.StatusBadRequest, message...)
}

func (e *HttpErrorUtil) Conflict(c *gin.Context, message ...string) {
	e.respondWithError(c, http.StatusConflict, message...)
}

func (e *HttpErrorUtil) NotFound(c *gin.Context, message ...string) {
	e.respondWithError(c, http.StatusNotFound, message...)
}

func (e *HttpErrorUtil) Unauthorized(c *gin.Context, message ...string) {
	e.respondWithError(c, http.StatusUnauthorized, message...)
}

func (e *HttpErrorUtil) Forbidden(c *gin.Context, message ...string) {
	e.respondWithError(c, http.StatusForbidden, message...)
}
