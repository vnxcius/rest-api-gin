package main

import (
	"github.com/vnxcius/gin-api/database"
	"github.com/vnxcius/gin-api/routes"
)

func main() {
	database.Connection()
	routes.HandleRequests()
}
