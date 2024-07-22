package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andiksetyawan/log/mocks"
	"github.com/andiksetyawan/response"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	mockLogger := &logmock.Logger{}
	data := map[string]string{"key": "value"}
	message := "Operation successful"
	expected := `{"status":"success","code":"OK","message":"Operation successful","data":{"key":"value"}}`

	t.Run("http.ResponseWriter", func(t *testing.T) {
		responder := response.NewResponder[http.ResponseWriter](mockLogger)

		w := httptest.NewRecorder()
		err := responder.Success(w, data, message)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("echo.Context", func(t *testing.T) {
		responder := response.NewResponder[echo.Context](mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := responder.Success(c, data, message)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("gin.Context", func(t *testing.T) {
		responder := response.NewResponder[*gin.Context](mockLogger)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		err := responder.Success(c, data, message)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})
}

func TestError(t *testing.T) {
	mockLogger := &logmock.Logger{}

	err := http.ErrBodyNotAllowed
	expected := `{"status":"error","code":"BAD_REQUEST","message":"Invalid input","errors":["http: request method or response status code does not allow body"]}`

	t.Run("http.ResponseWriter", func(t *testing.T) {
		responder := response.NewResponder[http.ResponseWriter](mockLogger)

		w := httptest.NewRecorder()
		err := responder.Error(w, http.StatusBadRequest, err, "Invalid input")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("echo.Context", func(t *testing.T) {
		responder := response.NewResponder[echo.Context](mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := responder.Error(c, http.StatusBadRequest, err, "Invalid input")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("gin.Context", func(t *testing.T) {
		responder := response.NewResponder[*gin.Context](mockLogger)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		err := responder.Error(c, http.StatusBadRequest, err, "Invalid input")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})
}

func TestErrorCustomCode(t *testing.T) {
	mockLogger := &logmock.Logger{}

	err := http.ErrBodyNotAllowed
	expected := `{"status":"error","code":"CUSTOM_ERROR_CODE","message":"Custom error","errors":["http: request method or response status code does not allow body"]}`

	t.Run("http.ResponseWriter", func(t *testing.T) {
		responder := response.NewResponder[http.ResponseWriter](mockLogger)

		w := httptest.NewRecorder()
		err := responder.ErrorCustomCode(w, http.StatusForbidden, "CUSTOM_ERROR_CODE", err, "Custom error")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("echo.Context", func(t *testing.T) {
		responder := response.NewResponder[echo.Context](mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := responder.ErrorCustomCode(c, http.StatusForbidden, "CUSTOM_ERROR_CODE", err, "Custom error")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("gin.Context", func(t *testing.T) {
		responder := response.NewResponder[*gin.Context](mockLogger)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		err := responder.ErrorCustomCode(c, http.StatusForbidden, "CUSTOM_ERROR_CODE", err, "Custom error")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})
}

func TestSuccessWithCode(t *testing.T) {
	mockLogger := &logmock.Logger{}

	data := map[string]string{"key": "value"}
	message := "Resource created"
	expected := `{"status":"success","code":"CREATED","message":"Resource created","data":{"key":"value"}}`

	t.Run("http.ResponseWriter", func(t *testing.T) {
		responder := response.NewResponder[http.ResponseWriter](mockLogger)

		w := httptest.NewRecorder()
		err := responder.SuccessWithCode(w, http.StatusCreated, data, message)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})

	t.Run("echo.Context", func(t *testing.T) {
		responder := response.NewResponder[echo.Context](mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := responder.SuccessWithCode(c, http.StatusCreated, data, message)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("gin.Context", func(t *testing.T) {
		responder := response.NewResponder[*gin.Context](mockLogger)

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		err := responder.SuccessWithCode(c, http.StatusCreated, data, message)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.JSONEq(t, expected, w.Body.String())
	})
}
