## Go Restapi Responder Package
### Example:
```
package main

import (
	"context"

	"github.com/andiksetyawan/log/slog"
	"github.com/andiksetyawan/response"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	logger, _ := slog.New()
	responder := response.NewResponder[echo.Context](logger)

	e.GET("/", func(c echo.Context) error {
		return responder.Success(c, nil, "Hello, World!")
	})

	err := e.Start(":2323")
	if err != nil {
		logger.Error(context.TODO(), "failed to run web server", "error", err)
	}
}

```
Ruquest:
```
curl --request GET --url http://localhost:2323/
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

### TODO
- [ ] http error wrap
- [ ] log error request header/body