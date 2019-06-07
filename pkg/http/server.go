package http

import (
	"cran/pkg/generating"
	"fmt"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yosssi/ace"
)

// Server represents the http server interface to host the application.
type Server interface {
	Start(port int) error
}

type server struct {
	service generating.Service
}

// NewServer instantiates a new http server to serve the application.
func NewServer(service generating.Service) Server {
	return &server{
		service: service,
	}
}

func (s *server) Start(port int) error {
	e := echo.New()
	e.Renderer = &aceRenderer{
		options: &ace.Options{
			DynamicReload: true,
			BaseDir:       "pkg/http/templates",
			FuncMap: template.FuncMap{
				"raw": func(source string) template.HTML {
					return template.HTML(source)
				},
				"add": func(a, b int) int {
					return a + b
				},
				"wrap": func(report *generating.Report, node generating.Node) *wrapped {
					return &wrapped{
						Report: report,
						Node:   node,
					}
				},
			},
		},
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("static", "pkg/http/static")

	e.GET("/", index)
	e.GET("/report", generate(s.service))

	return e.Start(fmt.Sprintf(":%d", port))
}

type wrapped struct {
	Report *generating.Report
	Node   generating.Node
}
