package main

import (
	"html/template"
	"os"

	cran "cran/domain"
	_ "cran/providers/assemblee_nationale"

	"github.com/kataras/iris"
)

type wrapedContext struct {
	Node   cran.Node
	Report *cran.Report
}

func main() {
	app := iris.Default()
	app.StaticServe("static", "/static")
	pug := iris.Pug("templates", ".pug")
	pug.Reload(true)
	app.RegisterView(pug)

	// Tiny func to return a raw HTML
	pug.AddFunc("raw", func(source string) template.HTML {
		return template.HTML(source)
	})

	pug.AddFunc("wrap", func(node cran.Node, report *cran.Report) *wrapedContext {
		return &wrapedContext{
			Node:   node,
			Report: report,
		}
	})

	// Get the index form
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.pug")
	})

	// Get an individual report
	app.Get("/report", func(ctx iris.Context) {
		src := ctx.URLParam("source")
		p, err := cran.GuessProvider(src)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		p.Fetch(src, func(report *cran.Report, err error) {
			ctx.View("report.pug", report)
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Run(iris.Addr(":" + port))
}
