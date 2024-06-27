package discourse

type Image struct {
	MaxWidth  int    `json:"max_width,omitempty"`
	MaxHeight int    `json:"max_height,omitempty"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	URL       string `json:"url"`
}
