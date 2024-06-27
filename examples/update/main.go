package main

import (
	"fmt"
	"os"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {

	site := os.Getenv("DISCOURSE_SITE")
	apiKey := os.Getenv("DISCOURSE_APIKEY")
	username := os.Getenv("DISCOURSE_USERNAME")
	discourseClient := discourse.NewClient(site, apiKey, username)

	// Get the first category and change color to orange
	allCategoriesResponse, err := discourse.ListCategories(discourseClient)

	if err != nil {
		fmt.Println(err)
	} else {
		firstCategory := allCategoriesResponse.CategoryList.Categories[0]
		_, err := discourse.UpdateCategoryByID(discourseClient, firstCategory.ID, &discourse.NewCategory{Color: "e95420"})

		if err != nil {
			fmt.Println("Category color update error:")
			fmt.Println(err)
		}

		// Get the first topic in the category and bookmark it
		firstCategoryContents, err := discourse.GetCategoryContentsBySlug(discourseClient, firstCategory.Slug)

		if err != nil {
			fmt.Println(err)
		} else {
			firstTopic := firstCategoryContents.TopicList.Topics[0]
			err := discourse.BookmarkTopicByID(discourseClient, firstTopic.ID)

			if err != nil {
				fmt.Println("Bookmark error:")
				fmt.Println(err)
			}
		}
	}
}
