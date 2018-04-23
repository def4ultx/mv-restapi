package models

// SearchResponse base type
type SearchResponse struct {
	Kind string  `json:"kind"`
	ETAG string  `json:"etag"`
	Item []Video `json:"items"`
}
