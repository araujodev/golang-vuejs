package router

import (
	"net/http"

	"github.com/araujodev/golang-vuejs/pkg/types/routes"
	HomeHandler "github.com/araujodev/golang-vuejs/src/controllers/home"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
