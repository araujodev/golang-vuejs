package app

import (
	"log"
	"net/http"

	"github.com/araujodev/golang-vuejs/src/system/router"

	"github.com/go-xorm/xorm"
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
	r.Init()
	http.ListenAndServe(s.port, r.Router)
}
