package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/andiksetyawan/log"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type responder[T any] struct {
	log log.Logger
}

// NewResponder creates a new responder instance
func NewResponder[T any](logger log.Logger) *responder[T] {
	return &responder[T]{log: logger}
}

func (r *responder[T]) respond(w T, httpStatusCode int, response interface{}) (err error) {
	switch writer := any(w).(type) {
	case http.ResponseWriter:
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(httpStatusCode)
		return json.NewEncoder(writer).Encode(response)
	case echo.Context:
		writer.Response().Header().Set("Content-Type", "application/json")
		return writer.JSON(httpStatusCode, response)
	case *gin.Context:
		writer.Header("Content-Type", "application/json")
		writer.JSON(httpStatusCode, response)
	default:
		panic("unsupported response writer type")
	}

	return
}

// Success writes a success response with status OK
func (r *responder[T]) Success(w T, data interface{}, msg string) error {
	return r.SuccessWithCode(w, http.StatusOK, data, msg)
}

// SuccessWithCode writes a success response with a custom status code
func (r *responder[T]) SuccessWithCode(w T, httpStatusCode int, data any, msg string) error {
	response := SuccessResponse[T]{
		Response: Response{
			Status:  "success",
			Code:    strings.ReplaceAll(strings.ToUpper(http.StatusText(httpStatusCode)), " ", "_"),
			Message: msg,
		},
		Data: data,
	}

	return r.respond(w, httpStatusCode, response)
}

// Error writes an error response
func (r *responder[T]) Error(w T, httpStatusCode int, err error, msg string) error {
	if msg == "" {
		msg = http.StatusText(httpStatusCode)
	}

	response := ErrorResponse{
		Response: Response{
			Status:  "error",
			Message: msg,
			Code:    strings.ReplaceAll(strings.ToUpper(http.StatusText(httpStatusCode)), " ", "_"),
		},
		Errors: strings.Split(err.Error(), "\n"),
	}
	return r.respond(w, httpStatusCode, response)
}

// ErrorCustomCode writes an error response with a custom error code
func (r *responder[T]) ErrorCustomCode(w T, httpStatusCode int, errorCode string, err error, msg string) error {
	response := ErrorResponse{
		Response: Response{
			Status:  "error",
			Message: msg,
			Code:    errorCode,
		},
		Errors: strings.Split(err.Error(), "\n"),
	}
	return r.respond(w, httpStatusCode, response)
}
