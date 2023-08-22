package main

import "github.com/knvi/modnevi/internal/server"

func main() {
	port := ":4000"

	if err := server.NewServer(port); err != nil {
		panic(err)
	}
}