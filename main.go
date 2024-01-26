package main

import (
	"context"
	"log"

	"github.com/atharv-bhadange/youtube-api-search/cmd"
	D "github.com/atharv-bhadange/youtube-api-search/db"
	U "github.com/atharv-bhadange/youtube-api-search/utils"
	"github.com/atharv-bhadange/youtube-api-search/yt"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	if godotenv.Load(".env") != nil {
		log.Fatalln("Error while loading environment variables")
	}

	// Initialize database
	err := D.Init()
	if err != nil {
		log.Fatalln("Error while initializing database", err)
	}

	defer D.Close()

	// Get api keys
	apiKeys, err := U.GetApiKeys()

	if err != nil {
		log.Fatalln("Error while getting api keys", err)
	}

	// get query from command line
	query := U.GetQuery()

	// start a goroutine to fetch videos from youtube api and insert into database
	go yt.GetVideos(query, apiKeys, context.Background())

	// Initialize fiber app
	app := cmd.InitApp()

	port, err := U.GetEnvValue("PORT")

	if err != nil {
		port = "8081"
	}

	// Start server
	app.Listen(":" + port)
}
