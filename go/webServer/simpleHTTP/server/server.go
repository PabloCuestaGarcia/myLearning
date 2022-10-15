package server

import "net/http"

func NewServer(addr string) *http.Server {

	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
