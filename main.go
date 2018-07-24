package main

import (
	"flag"
	"log"
	"os"

	"github.com/araujodev/golang-vuejs/src/system/app"
	DB "github.com/araujodev/golang-vuejs/src/system/db"
	"github.com/joho/godotenv"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Selecione a porta")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
	log.Println("PORT: " + port)
}

func main() {
	db, err := DB.Connect()
	if err != nil {
		panic(err)
	}

	s := app.NewServer()
	s.Init(port, db)
	s.Start()
}
