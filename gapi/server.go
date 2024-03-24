package api

import (
	"net/http"
	"omidh70/sweat-db/db"
	"omidh70/sweat-db/util"

	"github.com/urfave/negroni"
)

type Server struct {
	store  db.Store //### add it later
	router *negroni.Negroni
	config *util.Config
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {

	server := Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return &server, nil
}

func (server *Server) setupRouter() {

	//First routes Doesn't Need it
	// server.router.POST("/users", server.createUser)
	// server.router.POST("/users/login", server.loginUser)
	// server.router.POST("/users/renew_access", server.renewAccessToken)

	// //Below Handlers use MiddleWares
	// authRouteGroup := server.router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRouteGroup.POST("/accounts", server.createAccount)
	// authRouteGroup.GET("/accounts/:id", server.getAccount)

	// authRouteGroup.POST("/transfers", server.createTransfer)
}

func (server *Server) Start(addr string) error {
	return http.ListenAndServe(addr, server.router)
}
