package api

import "github.com/labstack/echo/v4"

type Route struct {
	Path        string
	Method      string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

type Router interface {
	GetRoutes() []Route
}
