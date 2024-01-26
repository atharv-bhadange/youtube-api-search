package services

import (
	D "github.com/atharv-bhadange/youtube-api-search/db"
	M "github.com/atharv-bhadange/youtube-api-search/models"
)

func GetVideos(offset int, limit int) ([]M.VideoDetails, int64, error) {

	var videos []M.VideoDetails
	var count int64

	// Get video count
	err := D.DbConn.Model(&M.VideoDetails{}).Count(&count).Error
	if err != nil {
		return videos, 0, err
	}

	// Get videos based on offset and limit
	err = D.DbConn.Order("published_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&videos).
		Error
	if err != nil {
		return videos, 0, err
	}

	return videos, count, nil

}
