package models

// Video base model
type Video struct {
	Kind    string  `json:"kind"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

// ID base model
type ID struct {
	VideoID   string `json:"videoid,omitempty"`
	ChannelID string `json:"channelid,omitempty"`
}

// Snippet base type
type Snippet struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ChannalTitle string    `json:"channeltitle"`
	Thumbnail    Thumbnail `json:"thumbnails"`
}

// Thumbnail base type
type Thumbnail struct {
	Default ThumbnailInfo `json:"default"`
	Medium  ThumbnailInfo `json:"medium"`
	High    ThumbnailInfo `json:"high"`
}

// ThumbnailInfo base type
type ThumbnailInfo struct {
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}
