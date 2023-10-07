package Handlers

import (
	"Silver/Silver/Helpers"
	"Silver/Silver/Structure"
	"Silver/Silver/Utils"
	"net/http"
)

type HttpHandler struct {
	Routes []Structure.Route
}

// This function should be rewritten soon. works for now tho
func (handler HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	for _, route := range handler.Routes {
		if route.Static {
			if route.FilePath == request.URL.Path {
				http.ServeFile(responseWriter, request, "./Views/"+route.FilePath)
			}

			continue
		}

		if route.Path.MatchString(request.URL.Path) {
			if string(route.Method) != request.Method {
				responseWriter.WriteHeader(http.StatusMethodNotAllowed)
				_, err := responseWriter.Write([]byte("405 Method Not Allowed"))

				if err != nil {
					panic("error!")
				}

				return
			}

			matches := route.Path.FindStringSubmatch(request.URL.Path)
			params := make(map[string]string)

			for i, name := range route.Path.SubexpNames() {
				if i > 0 && i <= len(matches) {
					params[name] = matches[i]
				}
			}

			req := Utils.Request{
				Params: params,
			}

			res := Utils.Response{}

			Helpers.CurrentRequest = request
			Helpers.CurrentResponseWriter = responseWriter

			route.Controller(req, res)
			return
		}
	}

}
