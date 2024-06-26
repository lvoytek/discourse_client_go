package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")
	categoryShowResponse, err := discourse.ShowCategory(discourseClient, 78)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Specific category:", categoryShowResponse.Category.Name)
	}

	allCategoriesResponse, err := discourse.ListCategories(discourseClient)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("First category in full list:", allCategoriesResponse.CategoryList.Categories[0].Name)
	}

}
