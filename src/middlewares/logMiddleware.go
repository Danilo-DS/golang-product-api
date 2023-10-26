package middlewares

import (
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Host, r.Method, r.RequestURI)
		nextFunc(w, r)
	}
}
