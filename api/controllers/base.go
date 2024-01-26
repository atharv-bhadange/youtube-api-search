package controllers

import (
	"math"
	"strconv"

	S "github.com/atharv-bhadange/youtube-api-search/api/services"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"ok": 1,
	})
}

func GetVideos(ctx *fiber.Ctx) error {

	// Get page and limit from query params
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid page number",
		})
	}
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid limit",
		})
	}

	// Validate page and limit
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get videos from database
	videos, videoCount, err := S.GetVideos(offset, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(videoCount) / float64(limit)))

	// Return response
	return ctx.JSON(fiber.Map{
		"data": fiber.Map{
			"videos": videos,
		},
		"currentPage":     page,
		"itemsPerPage":    limit,
		"totalItems":      videoCount,
		"totalPages":      totalPages,
		"hasNextPage":     page < totalPages,
		"hasPreviousPage": page > 1,
	})
}
