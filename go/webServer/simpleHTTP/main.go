package main

import "learning/server"

func main() {

	srv := server.NewServer(":8080")
	srv.ListenAndServe()

}
