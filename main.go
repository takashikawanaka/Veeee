package main

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
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
	fmt.Println("Server Start Up........")
	http.ListenAndServe(":8080", &Router{})
}

type Router struct{}

type Response struct {
	content_type string
	code         int
	bytes        []byte
}

func (r *Response) WriteResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", r.content_type)
	(*w).WriteHeader(r.code)
	(*w).Write(r.bytes)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	switch path[1] {
	case "":
		if r.Method == http.MethodGet {
			fmt.Printf("GET : %s\n", r.URL.Path)
			(&Response{
				content_type: "text/html;charset=utf-8",
				code:         http.StatusOK,
				bytes:        html,
			}).WriteResponse(&w)
			return
		}
	case "static":
		if r.Method == http.MethodGet && len(path) == 3 {
			fmt.Printf("GET : %s\n", r.URL.Path)
			res := SearchStatic(path[2])
			res.WriteResponse(&w)
			return
		}
	case "graphql":
		switch r.Method {
		case http.MethodGet:
			fmt.Println("GraphQL")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func SearchStatic(name string) Response {
	switch name {
	case "index.css":
		return Response{
			content_type: "text/css;charset=utf-8",
			code:         http.StatusOK,
			bytes:        css,
		}
	case "index.js":
		return Response{
			content_type: "text/javascript;charset=utf-8",
			code:         http.StatusOK,
			bytes:        js,
		}
	default:
		return Response{
			content_type: "text/plain;charset=utf-8",
			code:         http.StatusNotFound,
			bytes:        []byte("NotFound"),
		}
	}
}
