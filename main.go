package main

import (
	"net/http"

	"github.com/def4ultx/mv-restapi/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.BodyLimit("2M"),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:8080",
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderContentType,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				echo.GET,
				echo.POST,
			},
			MaxAge: 3600,
		}),
	)

	// health check
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// register services``
	services.Video(e.Group("/api"))

	// start server
	e.Logger.Fatal(e.Start(":3000"))
}
