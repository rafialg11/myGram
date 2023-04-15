package main

import (
	"myGram/database"
	"myGram/router"
	"os"
)

func main() {
	database.StartDB()
	var PORT = os.Getenv("PORT")
	r := router.StartApp()
	r.Run(":" + PORT)
}
