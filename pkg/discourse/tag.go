package discourse

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Count     int    `json:"count"`
	PMCount   int    `json:"pm_count"`
	TargetTag string `json:"target_tag"`
}

type TagData struct {
	Users         []PostCreator `json:"users"`
	PrimaryGroups []Group       `json:"primary_groups"`
	TopicList     TagTopicList  `json:"topic_list"`
}

type TagGroup struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	TagNames      []string       `json:"tag_names"`
	ParentTagName []string       `json:"parent_tag_name"`
	OnePerTopic   bool           `json:"one_per_topic"`
	Permissions   map[string]int `json:"permissions"`
}

type TagTopicList struct {
	CanCreateTopic bool              `json:"can_create_topic"`
	Draft          string            `json:"draft"`
	DraftKey       string            `json:"draft_key"`
	DraftSequence  int               `json:"draft_sequence"`
	PerPage        int               `json:"per_page"`
	Tags           []TagTopicListTag `json:"tags"`
	Topics         []SuggestedTopic  `json:"topics"`
}

type TagTopicListTag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TopicCount int    `json:"topic_count"`
	Staff      bool   `json:"staff"`
}

type ListTagsResponse struct {
	Tags   []Tag `json:"tags"`
	Extras struct {
		Categories []Category `json:"categories,omitempty"`
		TagGroups  []TagGroup `json:"tag_groups,omitempty"`
	} `json:"extras"`
}

type ListTagGroupsResponse struct {
	TagGroups []TagGroup `json:"tag_groups"`
}

func GetTagGroupByID(client *Client, id int) (response *TagGroup, err error) {
	data, sendErr := client.Get(fmt.Sprintf("tag_groups/%d", id), []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetTagByName(client *Client, name string) (response *TagData, err error) {
	data, sendErr := client.Get(fmt.Sprintf("tag/%s", name), []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func ListTags(client *Client) (response *ListTagsResponse, err error) {
	data, sendErr := client.Get("tags", []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func ListTagGroups(client *Client) (response *ListTagGroupsResponse, err error) {
	data, sendErr := client.Get("tag_groups", []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
