package http

import (
	"cran/pkg/generating"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.Render(200, "index", nil)
}

func generate(service generating.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		src := c.QueryParam("source")

		if src == "" {
			return echo.ErrBadRequest
		}

		service.Generate(src, func(r *generating.Report, err error) {
			if err != nil {
				c.Error(echo.ErrBadRequest.SetInternal(err))
				return
			}

			if err = c.Render(200, "report", r); err != nil {
				c.Logger().Error(err)
				c.Error(err)
			}
		})

		return nil
	}
}
