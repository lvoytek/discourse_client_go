package discourse

import (
	"testing"
	"time"
)

func TestGenerateSearchQueryString(t *testing.T) {
	tests := []struct {
		name  string
		query SearchQuery
		want  string
	}{
		{
			"All fields empty",
			SearchQuery{
				Username:         "",
				Category:         "",
				Tags:             []string{},
				RequiresAllTags:  false,
				Before:           time.Time{},
				After:            time.Time{},
				Order:            "",
				AssignedUsername: "",
				In:               []string{},
				With:             []string{},
				Status:           []string{},
				Group:            "",
				GroupMessages:    "",
				MinPosts:         0,
				MaxPosts:         -1,
				MinViews:         0,
				MaxViews:         -1,
			},
			"",
		},
		{
			"Defaults",
			SearchQuery{},
			"",
		},
		{
			"Time check",
			SearchQuery{
				Username:         "",
				Category:         "",
				Tags:             []string{},
				RequiresAllTags:  false,
				Before:           time.Date(2024, 4, 27, 10, 0, 0, 0, time.Local),
				After:            time.Date(2023, 10, 14, 18, 30, 0, 0, time.Local),
				Order:            "",
				AssignedUsername: "",
				In:               []string{},
				With:             []string{},
				Status:           []string{},
				Group:            "",
				GroupMessages:    "",
				MinPosts:         0,
				MaxPosts:         -1,
				MinViews:         0,
				MaxViews:         -1,
			},
			"before:2024-04-27 after:2023-10-14",
		},
		{
			"Tags In With and Status",
			SearchQuery{
				Username: "",
				Category: "",
				Tags: []string{
					"api",
					"docs",
					"test",
				},
				RequiresAllTags:  false,
				Before:           time.Time{},
				After:            time.Time{},
				Order:            "",
				AssignedUsername: "",
				In: []string{
					InBookmarks,
					InFirst,
				},
				With: []string{
					WithImages,
				},
				Status: []string{
					StatusClosed,
					StatusOpen,
				},
				Group:         "",
				GroupMessages: "",
				MinPosts:      0,
				MaxPosts:      -1,
				MinViews:      0,
				MaxViews:      -1,
			},
			"tags:api,docs,test in:bookmarks in:first with:images status:closed,open",
		},
		{
			"All tags required",
			SearchQuery{
				Username: "",
				Category: "",
				Tags: []string{
					"api",
					"docs",
					"test",
				},
				RequiresAllTags:  true,
				Before:           time.Time{},
				After:            time.Time{},
				Order:            "",
				AssignedUsername: "",
				In: []string{
					InBookmarks,
					InFirst,
				},
				With: []string{
					WithImages,
				},
				Status: []string{
					StatusClosed,
					StatusOpen,
				},
				Group:         "",
				GroupMessages: "",
				MinPosts:      0,
				MaxPosts:      -1,
				MinViews:      0,
				MaxViews:      -1,
			},
			"tags:api+docs+test in:bookmarks in:first with:images status:closed,open",
		},
		{
			"Other Fields",
			SearchQuery{
				Username:         "lvoytek",
				Category:         "server",
				Tags:             []string{},
				RequiresAllTags:  true,
				Before:           time.Time{},
				After:            time.Time{},
				Order:            OrderLatestTopic,
				AssignedUsername: "lvoytek",
				In:               []string{},
				With:             []string{},
				Status:           []string{},
				Group:            "canonical",
				GroupMessages:    "canonical",
				MinPosts:         3,
				MaxPosts:         -1,
				MinViews:         4,
				MaxViews:         -1,
			},
			"@lvoytek #server order:latest_topic assigned:lvoytek group:canonical group_messages:canonical min_posts:3 min_views:4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateSearchQueryString(&tt.query)
			if result != tt.want {
				t.Errorf("received \"%s\", want \"%s\"", result, tt.want)
			}
		})
	}
}
