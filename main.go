package main

import (
	"task-5-vix-btpns-SofyanEgiLesmana/database"
	"task-5-vix-btpns-SofyanEgiLesmana/router"
)

func main() {
	database.StartDB()
	database.MigateDB()
	router.SetupRouter().Run(":8080")
}
