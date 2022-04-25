package main

import (
	routes "websocket-boilerplate/route"
)

func main() {
	port := "8000"

	r := routes.SetupRouter()
	err := r.Run("localhost:" + port)
	if err != nil {
		return
	}
}
