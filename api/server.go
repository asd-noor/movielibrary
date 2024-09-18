package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	framework *echo.Echo
	// config    config.AppConfig
	routes []Route
}

func NewServer(routers ...Router) Server {
	rs := make([]Route, 0)
	for _, v := range routers {
		rs = append(rs, v.GetRoutes()...)
	}

	return Server{
		framework: echo.New(),
		routes:    rs,
	}
}

func (s *Server) Start() {
	s.registerRoutes()

	e := s.framework
	e.Logger.Fatal(e.Start(":" + "8080"))
}

func (s *Server) registerRoutes() {
	e := s.framework
	applyDefaultMiddlewares(e, defaultSkipper)

	for _, v := range s.routes {
		switch v.Method {
		case http.MethodGet:
			e.GET(v.Path, v.Handler, v.Middlewares...)
		case http.MethodPost:
			e.POST(v.Path, v.Handler, v.Middlewares...)
		case http.MethodPut:
			e.PUT(v.Path, v.Handler, v.Middlewares...)
		case http.MethodPatch:
			e.PATCH(v.Path, v.Handler, v.Middlewares...)
		case http.MethodDelete:
			e.DELETE(v.Path, v.Handler, v.Middlewares...)
		}
	}
}
