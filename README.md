## Go Restapi Responder Package
### Example:
```
package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/andiksetyawan/log/slog"
	"github.com/andiksetyawan/response"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	logger, _ := slog.New()
	responder, _ := response.NewResponder(response.WithErrLogger[echo.Context](logger, "response error"))

	e.GET("/success", func(c echo.Context) error {
		ctx := c.Request().Context()

		return responder.Success(ctx, c, nil, "Hello, World!")
	})

	e.POST("/success-with-code", func(c echo.Context) error {
		ctx := c.Request().Context()

		return responder.SuccessWithCode(ctx, c, http.StatusCreated, map[string]any{"id": 1}, "data has been successfully created")
	})

	e.POST("/err", func(c echo.Context) error {
		ctx := c.Request().Context()

		err := errors.New("field 'name' is required")
		return responder.Error(ctx, c, http.StatusBadRequest, err, "failed to submit data")
	})

	e.POST("/err-with-custom-code", func(c echo.Context) error {
		ctx := c.Request().Context()
		err := errors.New("field 'name' is required")
		customCode := "BAD_REQUEST_FIELD_REQUIRED"

		return responder.ErrorCustomCode(ctx, c, http.StatusBadRequest, customCode, err, "failed to submit data")
	})

	err := e.Start(":2323")
	if err != nil {
		logger.Error(context.TODO(), "failed to run web server", "error", err)
	}
}

```
Request:
```
curl --request GET --url http://localhost:2323/success
```
HTTP Response:
```
{
  "status": "success",
  "code": "OK",
  "message": "Hello, World!",
  "data": null
}
```
```
{
  "status": "success",
  "code": "CREATED",
  "message": "data has been successfully created",
  "data": {
    "id": 1
  }
}
```

```
{
  "status": "error",
  "code": "BAD_REQUEST",
  "message": "failed to submit data",
  "errors": [
    "field 'name' is required"
  ]
}
```
```
{
  "status": "error",
  "code": "BAD_REQUEST_FIELD_REQUIRED",
  "message": "failed to submit data",
  "errors": [
    "field 'name' is required"
  ]
}
```

### TODO
- [ ] http error wrap
- [ ] log error request header/body