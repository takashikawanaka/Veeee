package main

import (
	"embed"
	"log"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/takashikawanaka/Veeee/graph"
)

//go:embed web/dist/*
var web embed.FS

var html []byte = OpenFile("web/dist/index.html")
var css []byte = OpenFile("web/dist/index.css")
var js []byte = OpenFile("web/dist/index.js")

func OpenFile(name string) []byte {
	bytes, err := web.ReadFile(name)
	if err != nil {
		panic("Error")
	}
	return bytes
}

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		switch strings.Split(r.URL.Path, "/")[2] {
		case "index.html":
			http.Redirect(w, r, "/", http.StatusFound)
		case "index.css":
			http.ServeFile(w, r, "web/dist/index.css")
		case "index.js":
			http.ServeFile(w, r, "web/dist/index.js")
		default:
			http.NotFound(w, r)
		}
	})
	http.Handle("/playground/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query/", srv)
	log.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
