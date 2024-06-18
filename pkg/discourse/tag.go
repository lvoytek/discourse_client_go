package discourse

type Tag struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Count     int    `json:"count"`
	PMCount   int    `json:"pm_count"`
	TargetTag string `json:"target_tag"`
}

type TagGroup struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	TagNames      []string       `json:"tag_names"`
	ParentTagName []string       `json:"parent_tag_name"`
	OnePerTopic   bool           `json:"one_per_topic"`
	Permissions   map[string]int `json:"permissions"`
}

type ListTagsResponse struct {
	Tags   []Tag `json:"tags"`
	Extras struct {
		Categories []Category `json:"categories,omitempty"`
	} `json:"extras,omitempty"`
}
