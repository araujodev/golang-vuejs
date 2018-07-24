package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/araujodev/golang-vuejs/src/system/router"

	"github.com/go-xorm/xorm"

	"github.com/gorilla/handlers"
)

type Server struct {
	port string
	Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Iniciando Servidor ....")
	s.port = ":" + port
	s.Db = db
}

func (s *Server) Start() {
	log.Println("Començando servidor em " + s.port)

	r := router.NewRouter()
	r.Init(s.Db)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
}
