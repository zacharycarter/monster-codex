package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"

	"github.com/zacharycarter/monster-codex/config"
	"github.com/zacharycarter/monster-codex/db"
	"github.com/zacharycarter/monster-codex/handler"
)

var r *pat.Router

func init() {
	r = pat.New()

	r.Get("/api/monsters", handler.Monsters)
	r.Get("/", handler.Home)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/favicon.ico", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, "./static/images/favicon.png")
	})

	http.HandleFunc("/robots.txt", func(req http.ResponseWriter, res *http.Request) {
		http.ServeFile(req, res, "./static/robots.txt")
	})

	http.Handle("/", r)

	db.Connect()
	db.Initialize()
	defer db.Close()

	port := config.Port
	log.Println("Starting listen server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
