package main

import "github.com/gorilla/pat"

var router *pat.Router

func SetRouter(r *pat.Router) {
	router = r
}

func routeUrl(routeName string, urlArgs ...string) string {
	url, err := router.GetRoute(routeName).URL(urlArgs...)

	if err != nil {
		panic(err)
	}

	return url.String()
}
