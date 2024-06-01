package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")
	userBadgeResponse, err := discourse.ListBadgesForUser(discourseClient, "lvoytek")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userBadgeResponse)
	}

}
