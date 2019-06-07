package http

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/yosssi/ace"
)

type aceRenderer struct {
	options *ace.Options
}

func (r *aceRenderer) Render(writer io.Writer, name string, data interface{}, ctx echo.Context) error {

	tpl, err := ace.Load(name, "", r.options)

	if err != nil {
		return err
	}

	return tpl.Execute(writer, data)
}
