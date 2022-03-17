package server

import (
	"fmt"
	"net/http"

	"github.com/Dhliv/Go-Server/server/middleware"
	"github.com/Dhliv/Go-Server/server/router"
	"github.com/Dhliv/Go-Server/server/router/handlers"
)

type server struct {
	port   string
	router *router.Router
}

func NewServer(port string) *server {
	return &server{
		port:   port,
		router: router.NewRouter(),
	}
}

func (s *server) Listen() error {
	http.Handle("/", s.router)
	fmt.Println("Server running on http://localhost" + s.port)

	err := http.ListenAndServe(s.port, nil)
	return err
}

func (s *server) Handle(path, method string, handler http.HandlerFunc) {
	s.router.NewHandler(path, method, handler)
}

// Ejecuta todos los middlewares secuencialmente, estos middlewares son verificaciones que se pueden hacer.
// Al final se retorna la función que se desea ejecutar.
func (s *server) AddMiddleware(f http.HandlerFunc, middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, mwr := range middlewares {
		f = mwr(f)
	}

	return f
}

func Server() {
	ownServer := NewServer(":5000")
	// we add the routes to the servers
	ownServer.Handle("/", "GET", handlers.HandleRoot)
	ownServer.Handle("/api", "POST", ownServer.AddMiddleware(handlers.HandleHome, middleware.CheckAuth(), middleware.Loggin()))
	ownServer.Handle("/createUser", "POST", handlers.UserPostRequest)

	ownServer.Listen()
}
