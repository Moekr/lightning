package http

import (
	"html/template"
	"net/http"
	"os"

	"gopkg.in/macaron.v1"

	"github.com/Moekr/gopkg/logs"
	"github.com/Moekr/lightning/util/algo"
)

func StartHTTPService(address string) {
	m := macaron.New()
	if os.Getenv("LIGHTNING_DEV") != "" {
		m.Use(macaron.Logger())
	} else {
		macaron.Env = macaron.PROD
	}
	m.Use(macaron.Recovery())
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "tmpl",
		Funcs: []template.FuncMap{
			map[string]interface{}{
				"inc":  algo.Inc,
				"dec":  algo.Dec,
				"tags": ParseTags,
			},
		},
	}))
	m.Use(macaron.Static("assets", macaron.StaticOptions{
		Prefix:      "assets",
		SkipLogging: macaron.Env == macaron.PROD,
	}))

	m.Get("/", handleIndex)
	m.Get("/index.html", handleIndex)
	m.Get("/page/:name.html", handlePage)
	m.Get("/article/:name.html", handleArticle)
	m.Get("/article/:name.md", handleMarkdown)
	m.Get("/archive.html", handleArchive)
	m.Get("/search.html", handleSearch)
	m.NotFound(handleNotFound)
	m.InternalServerError(handleInternalServerError)

	go func() {
		if err := http.ListenAndServe(address, m); err != nil {
			panic(err)
		}
		logs.Info("[HTTP] service served on %s", address)
	}()
}
