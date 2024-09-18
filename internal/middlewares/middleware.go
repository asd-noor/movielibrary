package middlewares

import (
	"movielibrary/internal/domain/errors"
	"movielibrary/internal/utils"
	pc "movielibrary/pkg/consts"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthenticateAppKey(actualAppKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			appKey := c.Request().Header.Get(pc.AppKeyHeaderID)
			if actualAppKey != appKey {
				return c.JSON(http.StatusForbidden, utils.PrepareErrResponse(errors.ErrInvalidAppKey, ""))
			}

			return next(c)
		}
	}
}
