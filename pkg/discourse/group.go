package discourse

import (
	"encoding/json"
	"fmt"
)

type NewGroup struct {
	Name                            string `json:"name"`
	FullName                        string `json:"full_name"`
	BioRaw                          string `json:"bio_raw"`
	Usernames                       string `json:"usernames"`
	OwnerUsernames                  string `json:"owner_usernames"`
	AutomaticMembershipEmailDomains string `json:"automatic_membership_email_domains"`
	VisibilityLevel                 int    `json:"visibility_level"`
	PrimaryGroup                    bool   `json:"primary_group"`
	FlairIcon                       string `json:"flair_icon"`
	FlairUploadID                   int    `json:"flair_upload_id"`
	FlairBGColor                    string `json:"flair_bg_color"`
	PublicAdmission                 bool   `json:"public_admission"`
	PublicExit                      bool   `json:"public_exit"`
	DefaultNotificationLevel        int    `json:"default_notification_level"`
	MutedCategoryIDs                []int  `json:"muted_category_ids"`
	RegularCategoryIDs              []int  `json:"regular_category_ids"`
	WatchingCategoryIDs             []int  `json:"watching_category_ids"`
	TrackingCategoryIDs             []int  `json:"tracking_category_ids"`
	WatchingFirstPostCategoryIDs    []int  `json:"watching_first_post_category_ids"`
}

type Group struct {
	ID                        int    `json:"id"`
	Automatic                 bool   `json:"automatic"`
	Name                      string `json:"name"`
	UserCount                 int    `json:"user_count"`
	MentionableLevel          int    `json:"mentionable_level"`
	MessageableLevel          int    `json:"messageable_level"`
	VisibilityLevel           int    `json:"visibility_level"`
	PrimaryGroup              bool   `json:"primary_group"`
	Title                     string `json:"title"`
	GrantTrustLevel           int    `json:"grant_trust_level"`
	IncomingEmail             string `json:"incoming_email"`
	HasMessages               bool   `json:"has_messages"`
	FlairURL                  string `json:"flair_url"`
	FlairBGColor              string `json:"flair_bg_color"`
	FlairColor                string `json:"flair_color"`
	BioRaw                    string `json:"bio_raw"`
	BioCooked                 string `json:"bio_cooked"`
	BioExcerpt                string `json:"bio_excerpt"`
	PublicAdmission           bool   `json:"public_admission"`
	PublicExit                bool   `json:"public_exit"`
	AllowMembershipRequests   bool   `json:"allow_membership_requests"`
	FullName                  string `json:"full_name"`
	DefaultNotificationLevel  int    `json:"default_notification_level"`
	MembershipRequestTemplate string `json:"membership_request_template"`
	MembersVisibilityLevel    int    `json:"members_visibility_level"`
	CanSeeMembers             bool   `json:"can_see_members"`
	CanAdminGroup             bool   `json:"can_admin_group"`
	CanEditGroup              bool   `json:"can_edit_group"`
	PublishReadState          bool   `json:"publish_read_state"`
}

type GroupPermission struct {
	PermissionType int    `json:"permission_type"`
	GroupName      string `json:"group_name"`
}

type GroupUser struct {
	GroupID           int  `json:"group_id"`
	UserID            int  `json:"user_id"`
	NotificationLevel int  `json:"notification_level"`
	Owner             bool `json:"owner"`
}

type GetGroupResponse struct {
	Group Group `json:"group"`
}

type CreateGroupRequest struct {
	Group NewGroup `json:"group"`
}

type CreateGroupResponse struct {
	BasicGroup Group `json:"basic_group"`
}

type GroupMemberList struct {
	Members []User              `json:"members"`
	Owners  []User              `json:"owners"`
	Meta    GroupMemberMetadata `json:"meta"`
}

type GroupMemberMetadata struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func GetGroupByID(client *Client, id int) (response *GetGroupResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("groups/%d", id))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetGroupByName(client *Client, name string) (response *GetGroupResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("groups/%s", name))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetGroupMembersByID(client *Client, id int, offset int) (response *GroupMemberList, err error) {
	data, sendErr := client.GetWithQueryString(fmt.Sprintf("groups/%d/members", id), fmt.Sprintf("offset=%d", offset))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetGroupMembersByName(client *Client, name string, offset int) (response *GroupMemberList, err error) {
	data, sendErr := client.GetWithQueryString(fmt.Sprintf("groups/%s/members", name), fmt.Sprintf("offset=%d", offset))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func CreateGroup(client *Client, group *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	inputData, marshalError := json.Marshal(group)

	if marshalError != nil {
		return nil, marshalError
	}

	data, sendErr := client.PostWithReturn("admin/groups", inputData)

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func DeleteGroupByID(client *Client, id int) (err error) {
	return client.Delete(fmt.Sprintf("admin/groups/%d", id))
}
