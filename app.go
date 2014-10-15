package main

import (
	"github.com/gorilla/pat"
	"log"
	"net/http"
)

var r *pat.Router

func init() {
	r = pat.New()
	SetRouter(r)
	r.Get("/", NewHomeHandler(NewTemplateRenderer(), NewMonsterRepository()).Handle)
}

func main() {
	Connect(DatabaseConnectionString())
	defer Cleanup()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/favicon.ico", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, "./static/images/favicon.png")
	})

	http.HandleFunc("/robots.txt", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, "./static/robots.txt")
	})

	http.Handle("/", r)

	port := Port()
	log.Println("Starting listen server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
