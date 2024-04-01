package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"omidh70/sweat-db/db"
	"omidh70/sweat-db/gapi"
	"omidh70/sweat-db/pb"
	"omidh70/sweat-db/util"
	"time"

	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func _main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}

func main() {

	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg.DBHost = "localhost"
	cfg.DBPort = "27017"
	cfg.Timeout = 15 * time.Second

	store, err := db.NewMongoHandler(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	runGrpcServer(cfg, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(&config, store)
	if err != nil {
		log.Fatal("can't start grpc server")
	}

	// grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServer(grpcServer, server)
	pb.RegisterPingPongServer(grpcServer, server)
	//it allows user to explore self document
	reflection.Register(grpcServer)

	grpcListener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("can't start grpc listener")
	}
	log.Printf("start GRPC server on %s", config.GRPCServerAddress)

	err = grpcServer.Serve(grpcListener)
	if err != nil {
		log.Fatal("can't start server ", err)
	}
}
