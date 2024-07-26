package discourse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const (
	OrderLatest      string = "latest"
	OrderLikes       string = "likes"
	OrderViews       string = "views"
	OrderLatestTopic string = "latest_topic"
)

const (
	InTitle      string = "title"
	InLikes      string = "likes"
	InPersonal   string = "personal"
	InMessages   string = "messages"
	InSeen       string = "seen"
	InUnseen     string = "unseen"
	InPosted     string = "posted"
	InCreated    string = "created"
	InWatching   string = "watching"
	InTracking   string = "tracking"
	InBookmarks  string = "bookmarks"
	InAssigned   string = "assigned"
	InUnassigned string = "unassigned"
	InFirst      string = "first"
	InPinned     string = "pinned"
	InWiki       string = "wiki"
)

const (
	WithImages string = "images"
)

const (
	StatusOpen       string = "open"
	StatusClosed     string = "closed"
	StatusPublic     string = "public"
	StatusArchived   string = "archived"
	StatusNoReplies  string = "noreplies"
	StatusSingleUser string = "single_user"
	StatusSolved     string = "solved"
	StatusUnsolved   string = "unsolved"
)

type SearchQuery struct {
	Term             string
	Username         string
	Category         string
	Tags             []string
	RequiresAllTags  bool
	Before           time.Time
	After            time.Time
	Order            string
	AssignedUsername string
	In               []string
	With             []string
	Status           []string
	Group            string
	GroupMessages    string
	MinPosts         int
	MaxPosts         int
	MinViews         int
	MaxViews         int
	Custom           map[string][]string
}

type SearchResult struct {
	Posts               []PostData          `json:"posts"`
	Users               []User              `json:"users"`
	Categories          []Category          `json:"categories"`
	Tags                []TagData           `json:"tags"`
	Groups              []Group             `json:"groups"`
	GroupedSearchResult GroupedSearchResult `json:"grouped_search_result"`
}

type GroupedSearchResult struct {
	MorePosts           string                 `json:"more_posts"`
	MoreUsers           string                 `json:"more_users"`
	MoreCategories      string                 `json:"more_categories"`
	Term                string                 `json:"term"`
	SearchLogID         int                    `json:"search_log_id"`
	MoreFullPageResults string                 `json:"more_full_page_results"`
	CanCreateTopic      bool                   `json:"can_create_topic"`
	Error               string                 `json:"error"`
	Extra               map[string]interface{} `json:"extra,omitempty"`
	PostIDs             []int                  `json:"post_ids"`
	UserIDs             []int                  `json:"user_ids"`
	CategoryIDs         []int                  `json:"category_ids"`
	TagIDs              []int                  `json:"tag_ids"`
	GroupIDs            []int                  `json:"group_ids"`
}

func Search(client *Client, query *SearchQuery) (response *SearchResult, err error) {

	data, sendErr := client.GetWithQueryString("search", url.QueryEscape(generateSearchQueryString(query)))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func generateSearchQueryString(query *SearchQuery) string {
	searchString := "q="

	if query.Term != "" {
		searchString = fmt.Sprintf("%s%s ", searchString, query.Term)
	}

	if query.Username != "" {
		searchString = fmt.Sprintf("%s@%s ", searchString, query.Username)
	}

	if query.Category != "" {
		searchString = fmt.Sprintf("%s#%s ", searchString, query.Category)
	}

	tagString := ""

	for index, tagName := range query.Tags {
		tagString += tagName

		if index < len(query.Tags)-1 {
			if query.RequiresAllTags {
				tagString += "+"
			} else {
				tagString += ","
			}
		}
	}

	if tagString != "" {
		searchString = fmt.Sprintf("%stags:%s ", searchString, tagString)
	}

	if !query.Before.IsZero() {
		searchString = fmt.Sprintf("%sbefore:%d-%02d-%02d ", searchString, query.Before.Year(), query.Before.Month(), query.Before.Day())
	}

	if !query.After.IsZero() {
		searchString = fmt.Sprintf("%safter:%d-%02d-%02d ", searchString, query.After.Year(), query.After.Month(), query.After.Day())
	}

	if query.Order != "" {
		searchString = fmt.Sprintf("%sorder:%s ", searchString, query.Order)
	}

	if query.AssignedUsername != "" {
		searchString = fmt.Sprintf("%sassigned:%s ", searchString, query.AssignedUsername)
	}

	for _, inItem := range query.In {
		searchString = fmt.Sprintf("%sin:%s ", searchString, inItem)
	}

	for _, withItem := range query.With {
		searchString = fmt.Sprintf("%swith:%s ", searchString, withItem)
	}

	statusString := ""

	for index, statusName := range query.Status {
		statusString += statusName

		if index < len(query.Status)-1 {
			statusString += ","
		}
	}

	if statusString != "" {
		searchString = fmt.Sprintf("%sstatus:%s ", searchString, statusString)
	}

	if query.Group != "" {
		searchString = fmt.Sprintf("%sgroup:%s ", searchString, query.Group)
	}

	if query.GroupMessages != "" {
		searchString = fmt.Sprintf("%sgroup_messages:%s ", searchString, query.GroupMessages)
	}

	if query.MinPosts > 0 {
		searchString = fmt.Sprintf("%smin_posts:%d ", searchString, query.MinPosts)
	}

	if query.MaxPosts > 0 {
		searchString = fmt.Sprintf("%smax_posts:%d ", searchString, query.MaxPosts)
	}

	if query.MinViews > 0 {
		searchString = fmt.Sprintf("%smin_views:%d ", searchString, query.MinViews)
	}

	if query.MaxViews > 0 {
		searchString = fmt.Sprintf("%smax_views:%d ", searchString, query.MaxViews)
	}

	for customField, customData := range query.Custom {
		customDataString := ""

		for index, customItem := range customData {
			customDataString += customItem

			if index < len(customData)-1 {
				customDataString += ","
			}
		}

		if customDataString != "" {
			searchString = fmt.Sprintf("%s%s:%s ", searchString, customField, customDataString)
		}
	}

	if searchString != "" {
		searchString = searchString[:len(searchString)-1]
	}

	if searchString == "q" {
		searchString = ""
	}

	return searchString
}
