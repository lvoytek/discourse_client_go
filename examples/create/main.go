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

	// Create a new category on the site
	newCategoryData := discourse.NewCategory{
		Name:                      "New test category",
		Color:                     "31184f",
		TextColor:                 "ebebeb",
		AllowBadges:               true,
		Slug:                      "A test category made through the API",
		TopicFeaturedLinksAllowed: true,
		Permissions:               map[string]int{"everyone": 1},
	}

	newCategory, err := discourse.CreateCategory(discourseClient, &newCategoryData)

	if err != nil {
		fmt.Println(err)
	} else {
		// Create a new topic in the new category
		newCategoryID := newCategory.Category.ID
		newTopicData := discourse.NewPost{
			Title:    "New test topic title",
			Raw:      "This is the main content for a new api topic.",
			Category: newCategoryID,
		}

		newTopic, err := discourse.CreateTopic(discourseClient, &newTopicData)

		if err != nil {
			fmt.Println(err)
		} else {
			newTopicID := newTopic.TopicID
			newPostData := discourse.NewPost{
				Title:   "New post in topic",
				Raw:     "This is a new post in the new topic.",
				TopicID: newTopicID,
			}

			newPost, err := discourse.CreatePost(discourseClient, &newPostData)

			if err != nil {
				fmt.Println(err)
			} else {
				newPostID := newPost.ID
				fmt.Println("Created category", newCategoryID, "topic", newTopicID, "post", newPostID)
			}

		}
	}
}
