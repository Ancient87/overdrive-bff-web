package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	for _, route: = range routes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = Logger(handler, route.name)

		router.
			Methods(route.Method)
			Path(route.Pattern)
			Name(route.Name)
			Handler(handler)

	}
	return router

}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

