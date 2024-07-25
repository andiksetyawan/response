package response

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/andiksetyawan/log"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type responder[T any] struct {
	log       log.Logger
	logAttrs  []any
	logErrMsg string
}

type OptFunc[T any] func(o *responder[T]) error

func WithErrLogger[T any](logger log.Logger, logErrMsg string, attrs ...any) OptFunc[T] {
	return func(h *responder[T]) (err error) {
		h.log = logger
		h.logAttrs = append(h.logAttrs, attrs...)
		h.logErrMsg = logErrMsg
		return
	}
}

// NewResponder creates a newriterresponder instance
func NewResponder[T any](opt ...OptFunc[T]) (h *responder[T], err error) {
	h = new(responder[T])

	for _, fn := range opt {
		err = fn(h)
		if err != nil {
			return
		}
	}

	return
}

func (r *responder[T]) respond(writer T, httpStatusCode int, response interface{}) (err error) {
	switch writer := any(writer).(type) {
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

func (r *responder[T]) logErr(ctx context.Context, err error) {
	var attrs []any
	if len(r.logAttrs) != 0 {
		attrs = append(attrs, r.logAttrs)
	}

	message := r.logErrMsg
	if message == "" {
		message = "response error"
	}

	attrs = append(attrs, "error", err)
	r.log.Error(ctx, message, attrs...)
}

// Success writes a success response with status OK
func (r *responder[T]) Success(ctx context.Context, writer T, data interface{}, msg string) error {
	return r.SuccessWithCode(ctx, writer, http.StatusOK, data, msg)
}

// SuccessWithCode writes a success response with a custom status code
func (r *responder[T]) SuccessWithCode(ctx context.Context, writer T, httpStatusCode int, data any, msg string) error {
	response := SuccessResponse[T]{
		Response: Response{
			Status:  "success",
			Code:    strings.ReplaceAll(strings.ToUpper(http.StatusText(httpStatusCode)), " ", "_"),
			Message: msg,
		},
		Data: data,
	}

	return r.respond(writer, httpStatusCode, response)
}

// Error writes an error response
func (r *responder[T]) Error(ctx context.Context, writer T, httpStatusCode int, err error, msg string) error {
	if msg == "" {
		msg = http.StatusText(httpStatusCode)
	}

	if r.logErrMsg != "" {
		r.logErr(ctx, err)
	}

	response := ErrorResponse{
		Response: Response{
			Status:  "error",
			Message: msg,
			Code:    strings.ReplaceAll(strings.ToUpper(http.StatusText(httpStatusCode)), " ", "_"),
		},
		Errors: strings.Split(err.Error(), "\n"),
	}
	return r.respond(writer, httpStatusCode, response)
}

// ErrorCustomCode writes an error response with a custom error code
func (r *responder[T]) ErrorCustomCode(ctx context.Context, writer T, httpStatusCode int, errorCode string, err error, msg string) error {
	response := ErrorResponse{
		Response: Response{
			Status:  "error",
			Message: msg,
			Code:    errorCode,
		},
		Errors: strings.Split(err.Error(), "\n"),
	}
	return r.respond(writer, httpStatusCode, response)
}
