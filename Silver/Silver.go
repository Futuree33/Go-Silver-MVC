package Silver

import (
	"Silver/Silver/Handlers"
	"Silver/Silver/Helpers"
	"Silver/Silver/Structure"
	"log"
	"net/http"
	"os"
	"strconv"
)

type App struct {
	port   int
	routes []Structure.Route
}

const (
	Get  Structure.RequestMethod = "GET"
	Post Structure.RequestMethod = "POST"
)

func (app *App) AddRoute(method Structure.RequestMethod, path string, controller Structure.Controller) {
	route := Structure.Route{
		Path:       Helpers.Regex(path),
		Controller: controller,
		Method:     method,
		Static:     false,
	}

	app.routes = append(app.routes, route)
}

func (app *App) AddStatic(dirName string) {
	dir, err := os.ReadDir("./Views/" + dirName)

	if err != nil {
		log.Fatal("Failed to read dir: " + err.Error())
	}

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		route := Structure.Route{
			FilePath: "/" + dirName + "/" + entry.Name(),
			Method:   Get,
			Static:   true,
		}

		app.routes = append(app.routes, route)
	}
}

func (app *App) Listen(port int) {
	app.port = port
	handler := Handlers.HttpHandler{
		Routes: app.routes,
	}

	println("http://localhost:" + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}
