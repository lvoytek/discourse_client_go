package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	const topicToGet = 22706

	discourseClient := discourse.NewAnonymousClient("https://meta.discourse.org")
	GetTopicResponse, err := discourse.GetTopicByID(discourseClient, topicToGet)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Topic #", topicToGet, "-", GetTopicResponse.Title)

		creatorName := GetTopicResponse.Details.CreatedBy.Name
		if creatorName == "" {
			creatorName = GetTopicResponse.Details.CreatedBy.Username
		}

		fmt.Println("Created By:", creatorName)
		fmt.Println("Created at", GetTopicResponse.CreatedAt)
	}

}
