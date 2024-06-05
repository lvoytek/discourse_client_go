package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")
	categoryShowResponse, err := discourse.ShowCategory(discourseClient, 12)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(categoryShowResponse.Category.Name)
	}

}
