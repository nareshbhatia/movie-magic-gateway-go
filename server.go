package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph"
	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph/resolver"
	moviepb "github.com/nareshbhatia/movie-magic-services-go/gen/go/movie/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "8080"

var (
	addr = flag.String("addr", "localhost:30000", "the address to connect to")
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	movieServiceClient := moviepb.NewMovieServiceClient(conn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver.NewResolver(movieServiceClient)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
