package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")

	// Get the latest post across topics
	latestPosts, err := discourse.GetLatestPosts(discourseClient)

	if err != nil {
		fmt.Println(err)
	} else {
		latestPostInfo := latestPosts.LatestPosts[0]
		fmt.Println("Latest post by", latestPostInfo.DisplayUsername)
		fmt.Println(latestPostInfo.Raw)

	}

	fmt.Println()

	// Get a specific post and its latest reply
	specificPost, err := discourse.GetPostByID(discourseClient, 77484)

	if err != nil {
		fmt.Println(err)
	} else {
		specificPostReplies, err := discourse.GetPostRepliesByID(discourseClient, specificPost.ID)

		if err != nil {
			fmt.Println(err)
		} else {
			latestReply := specificPostReplies[len(specificPostReplies)-1]
			fmt.Println("Latest reply to", specificPost.DisplayUsername, "by", latestReply.DisplayUsername)

			latestReplyFullData, err := discourse.GetPostByID(discourseClient, latestReply.ID)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(latestReplyFullData.Raw)
			}
		}
	}

	fmt.Println()

	// Get the latest post revision info for a specific post
	const revisedPostID = 77485

	numPostRevisions, err := discourse.GetNumPostRevisionsByID(discourseClient, revisedPostID)

	if err != nil {
		fmt.Println(err)
	} else {
		latestPostRevision, err := discourse.GetPostRevisionByID(discourseClient, revisedPostID, numPostRevisions)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Latest revision of", revisedPostID, "created by", latestPostRevision.Username, "- version", latestPostRevision.CurrentRevision)
		}
	}
}
