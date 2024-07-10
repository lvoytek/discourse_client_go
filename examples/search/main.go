package main

import (
	"fmt"
	"time"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")

	// Search for posts by lvoytek created between 2023-10-14 and 2024-04-27
	const timeLayout = "2006-01-02"

	before, _ := time.Parse(timeLayout, "2024-04-27")
	after, _ := time.Parse(timeLayout, "2023-10-14")

	username := "lvoytek"

	searchQuery := discourse.SearchQuery{
		Username: username,
		Before:   before,
		After:    after,
	}

	results, err := discourse.Search(discourseClient, &searchQuery)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Posts by", username, "created between", after, "and", before)

		for _, post := range results.Posts {
			fmt.Println(post.ID, "@", post.CreatedAt)
		}
	}
}
