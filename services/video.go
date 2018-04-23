package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/def4ultx/mv-restapi/api"
	"github.com/def4ultx/mv-restapi/models"
	"github.com/labstack/echo"
)

type requestBody struct {
	Title        string `json:"title"`
	ChannelTitle string `json:"channel_title"`
}

// type responseBody struct {
// 	Kind      string           `json:"kind"`
// 	ETAG      string           `json:"etag"`
// 	Item      Video            `json:"items"`
// 	Thumbnail models.Thumbnail `json:"thumbnail"`
// }

// Video init
func Video(g *echo.Group) {
	g.GET("/video", searchVideoHandler)
	g.POST("/video", searchByTitleHandler)
	g.GET("/video/:id", searchByIDHandler)
}

func searchVideoHandler(c echo.Context) error {
	metadata, err := api.RequestVideo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error.")
	}
	return c.JSON(http.StatusOK, metadata)
}

func searchByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// var body requestBody
	// if err := c.Bind(&body); err != nil {
	// 	return c.String(http.StatusBadRequest, "Bad request")
	// }
	// video, err := api.GetVideoByID(body.ID)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, "Internal server error.")
	// }
	// return c.JSON(http.StatusOK, video)
	fmt.Println(id)
	var video *models.Video
	if video = api.GetVideoByID(id); video == nil {
		return c.String(http.StatusInternalServerError, "Internal server error.")
	}
	return c.JSON(http.StatusOK, video)
}

func searchByTitleHandler(c echo.Context) error {
	var body requestBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}
	metadata, err := api.RequestVideo()
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error.")
	}
	var items []models.Video
	var name string
	if body.Title != "" {
		name = strings.ToLower(body.Title)
	} else {
		name = strings.ToLower(body.ChannelTitle)
	}
	for _, v := range metadata.Item {
		title := strings.ToLower(v.Snippet.Title)
		if strings.Contains(title, name) {
			items = append(items, v)
		}
	}
	return c.JSON(http.StatusOK, models.SearchResponse{
		Kind: metadata.Kind,
		ETAG: metadata.ETAG,
		Item: items,
	})
}
