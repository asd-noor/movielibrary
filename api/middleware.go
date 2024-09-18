package api

import (
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

func applyDefaultMiddlewares(e *echo.Echo, skipper em.Skipper) {
	e.Pre(em.RemoveTrailingSlash())

	e.Use(em.LoggerWithConfig(em.LoggerConfig{
		Skipper: skipper,
		Format: `{"time":"${time_custom}",` +
			`"host":"${host}","method":"${method}",` +
			`"status":${status},"error":"${error}"` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	e.Use(em.CORS(), em.Secure(), em.Recover())

	e.Use(em.GzipWithConfig(em.GzipConfig{
		Skipper: func(context echo.Context) bool {
			return context.Request().URL.Path == "/metrics"
		},
		Level: 5,
	}))
}

func defaultSkipper(_ echo.Context) bool {
	return false
}
