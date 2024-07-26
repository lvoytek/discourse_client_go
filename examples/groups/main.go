package main

import (
	"fmt"

	"github.com/lvoytek/discourse_client_go/pkg/discourse"
)

func main() {
	discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")

	const groupName = "canonical"

	groupResponse, err := discourse.GetGroupByName(discourseClient, groupName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number of users in", groupResponse.Group.FullName, "-", groupResponse.Group.UserCount)
	}

	membersResponse, err := discourse.GetGroupMembersByName(discourseClient, groupName, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("First member in group:", membersResponse.Members[0].Username)
	}

	fiftiesMemberResponse, err := discourse.GetGroupMembersByName(discourseClient, groupName, 50)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Fiftieth member in group:", fiftiesMemberResponse.Members[0].Username)
	}
}
