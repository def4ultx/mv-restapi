package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/def4ultx/mv-restapi/models"
)

const httpURI = "https://s3-ap-southeast-1.amazonaws.com/ysetter/media/video-search.json"

// RequestVideo get video metadata from httpURI and return SearchResponse
func RequestVideo() (*models.SearchResponse, error) {
	var body string
	res, err := http.Get(httpURI)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		body = string(bodyBytes)
	}
	// fmt.Println(body)
	// io.Copy(os.Stdout, res.Body)
	metadata := &models.SearchResponse{}
	err = json.Unmarshal([]byte(body), metadata)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return metadata, nil
	// video := &models.Video{}
	// err = json.Unmarshal([]byte(metadata.Item[0]), video)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(metadata.Item[0].Snippet.Title)
}

// GetVideoByID find video by ETAG and return
func GetVideoByID(id string) *models.Video {
	metadata, err := RequestVideo()
	if err != nil {
		return nil
	}
	for _, v := range metadata.Item {
		if v.ID.VideoID == id {
			video := v
			return &video
		}
	}
	return nil
}

// GetVideoByTitle find video by ETAG and return
// func GetVideoByTitle(title string) []models.Video {
// 	metadata, err := RequestVideo()
// 	if err != nil {
// 		return nil
// 	}
// 	var items []models.Video
// 	for _, v := range metadata.Item {
// 		title := strings.ToLower(v.Snippet.Title)
// 		if strings.Contains(title, title) {
// 			items = append(items, v)
// 		}
// 	}
// 	return items
// }
