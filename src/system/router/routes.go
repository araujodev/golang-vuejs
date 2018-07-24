package router

import (
	"net/http"

	"github.com/go-xorm/xorm"

	"github.com/araujodev/golang-vuejs/pkg/types/routes"
	AuthHandler "github.com/araujodev/golang-vuejs/src/controllers/auth"
	HomeHandler "github.com/araujodev/golang-vuejs/src/controllers/home"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	HomeHandler.Init(db)
	AuthHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
