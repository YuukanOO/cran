package main

import (
	"cran/cran"
	"html/template"
	"os"

	"github.com/kataras/iris"
)

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
