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

	allCategoriesResponse, err := discourse.ListCategories(discourseClient, true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("First category in full list:", allCategoriesResponse.CategoryList.Categories[0].Name)

		for _, nextCategory := range allCategoriesResponse.CategoryList.Categories {
			if len(nextCategory.SubcategoryIDs) > 0 {
				fmt.Println("First category in list with subcategories:", nextCategory.Name)
				fmt.Println("Subcategories include:")
				for _, nextSubcategory := range nextCategory.SubcategoryList {
					fmt.Println("\t", nextSubcategory.Name)
				}
				break
			}
		}
	}

	categorySecondPage, err := discourse.GetCategoryContentsBySlug(discourseClient, "server", 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("First server topic on second page:", categorySecondPage.TopicList.Topics[0].Title)
	}
}
