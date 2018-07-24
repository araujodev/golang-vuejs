package router

import (
	"net/http"

	"github.com/araujodev/golang-vuejs/pkg/types/routes"

	StatusHandler "github.com/araujodev/golang-vuejs/src/controllers/v1/status"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Token Get From HTTP Headers
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		//#END Token Get From HTTP Headers

		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {

	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}
	return
}
