package gapi

import (
	"context"
	"net/http"
	"omidh70/sweat-db/db"
	"omidh70/sweat-db/pb"
	"omidh70/sweat-db/util"

	"github.com/urfave/negroni"
)

type Server struct {
	store  db.Store //### add it later
	router *negroni.Negroni
	config *util.Config
	pb.UnimplementedGreeterServer
	pb.UnimplementedPingPongServer
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {

	server := Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return &server, nil
}

func (server *Server) SayHello(context.Context, *pb.CreateSweatRequest) (*pb.CreateSweatResponse, error) {
	return nil, nil
}

func (server *Server) SayBye(context.Context, *pb.CreateSweatRequest) (*pb.CreateSweatResponse, error) {
	return nil, nil
}

func (server *Server) setupRouter() {
	//TODO
}

func (server *Server) Start(addr string) error {
	return http.ListenAndServe(addr, server.router)
}
