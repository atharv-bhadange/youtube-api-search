package main

import (
	"log"

	"github.com/atharv-bhadange/youtube-api-search/cmd"
	"github.com/atharv-bhadange/youtube-api-search/db"
	"github.com/atharv-bhadange/youtube-api-search/utils"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	if godotenv.Load(".env") != nil {
		log.Fatalln("Error while loading environment variables")
	}

	// Initialize database
	err := db.Init()
	if err != nil {
		log.Fatalln("Error while initializing database", err)
	}

	defer db.Close()

	// Initialize fiber app
	app := cmd.InitApp()

	port, err := utils.GetEnvValue("PORT")

	if err != nil {
		port = "8081"
	}

	// Start server
	app.Listen("localhost:"+port)
}
