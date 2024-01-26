package yt

import (
	"context"
	"log"
	"time"

	"github.com/atharv-bhadange/youtube-api-search/constants"
	D "github.com/atharv-bhadange/youtube-api-search/db"
	M "github.com/atharv-bhadange/youtube-api-search/models"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetVideos(query string, api_keys []string, ctx context.Context) {

	parts := []string{"id", "snippet"}

	// Create a new YouTube client
	client, err := youtube.NewService(ctx, option.WithAPIKey(api_keys[0]))
	if err != nil {
		log.Fatalf("Failed to create YouTube API client: %v", err)
	}

	countRetries := 0

	for {

		stringTime := time.Now().Format(time.RFC3339)

		// Call the YouTube API to list videos for query cat
		resp, err := client.Search.List(parts).
			Q(query).
			Type("video").
			PublishedAfter(stringTime).
			MaxResults(50).
			Order("date").
			Do()
		if err != nil {
			log.Printf("Failed to search for videos: %v", err)
		}

		// retry if api cooldown
		if len(resp.Items) == 0 {
			log.Println("API cooldown, retrying...")

			countRetries++
			if countRetries < constants.MAX_RETRY_COUNT {
				time.Sleep(5 * time.Second)
				continue
			} else {
				break
			}
		}

		countRetries = 0

		// insert into database
		for _, item := range resp.Items {

			publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				log.Printf("Failed to parse publishedAt: %v", err)
			}

			video := M.VideoDetails{
				Id:          item.Id.VideoId,
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: publishedAt,
				Thumbnail:   item.Snippet.Thumbnails.Default.Url,
				Query:       query,
			}

			// check if video already exists
			var count int64
			D.DbConn.Model(&M.VideoDetails{}).Where("id = ?", video.Id).Count(&count)

			if count == 0 {
				D.DbConn.Create(&video)
			}

			// sleep for 500ms between each insert
			time.Sleep(500 * time.Millisecond)
		}

		// sleep for 10s between each api call
		time.Sleep(10 * time.Second)
	}
}
