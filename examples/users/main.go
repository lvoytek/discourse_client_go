package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")
	userResponse, err := discourse.GetUserByUsername(discourseClient, "lvoytek")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User - id:", userResponse.User.ID, "name:", userResponse.User.Name)
	}
}
