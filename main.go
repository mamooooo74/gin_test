package main

import "gin_test/server"

func main() {
	router := server.GetRouter()
	router.Run(":8080")
}
