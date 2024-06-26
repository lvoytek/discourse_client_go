package discourse

import (
	"encoding/json"
	"fmt"
)

type NewCategory struct {
	Name                      string         `json:"name"`
	Color                     string         `json:"color"`
	TextColor                 string         `json:"text_color"`
	ParentCategoryID          int            `json:"parent_category_id,omitempty"`
	AllowBadges               bool           `json:"allow_badges"`
	Slug                      string         `json:"slug"`
	TopicFeaturedLinksAllowed bool           `json:"topic_featured_links_allowed"`
	Permissions               map[string]int `json:"permissions"`
	SearchPriority            int            `json:"search_priority"`
	FormTemplateIDs           map[string]int `json:"form_template_ids"`
}

type Category struct {
	ID                           int               `json:"id"`
	Name                         string            `json:"name"`
	Color                        string            `json:"color"`
	TextColor                    string            `json:"text_color"`
	Slug                         string            `json:"slug"`
	TopicCount                   int               `json:"topic_count"`
	PostCount                    int               `json:"post_count"`
	Position                     int               `json:"position"`
	Description                  string            `json:"description"`
	DescriptionText              string            `json:"description_text"`
	DescriptionExcerpt           string            `json:"description_excerpt"`
	TopicURL                     string            `json:"topic_url"`
	ReadRestricted               bool              `json:"read_restricted"`
	Permission                   int               `json:"permission"`
	NotificationLevel            int               `json:"notification_level"`
	CanEdit                      bool              `json:"can_edit"`
	TopicTemplate                string            `json:"topic_template"`
	FormTemplateIDs              []int             `json:"form_template_ids"`
	HasChildren                  bool              `json:"has_children"`
	SubcategoryCount             int               `json:"subcategory_count"`
	SortOrder                    string            `json:"sort_order"`
	SortAscending                string            `json:"sort_ascending"`
	ShowSubcategoryList          bool              `json:"show_subcategory_list"`
	NumFeaturedTopics            int               `json:"num_featured_topics"`
	DefaultView                  string            `json:"default_view"`
	SubcategoryListStyle         string            `json:"subcategory_list_style"`
	DefaultTopPeriod             string            `json:"default_top_period"`
	DefaultListFilter            string            `json:"default_list_filter"`
	MinimumRequiredTags          int               `json:"minimum_required_tags"`
	NavigateToFirstPostAfterRead bool              `json:"navigate_to_first_post_after_read"`
	CustomFields                 map[string]string `json:"custom_fields"`
	AllowedTags                  []Tag             `json:"allowed_tags"`
	AllowedTagGroups             []TagGroup        `json:"allowed_tag_groups"`
	AllowGlobalTags              bool              `json:"allow_global_tags"`
	RequiredTagGroups            []struct {
		Name     string `json:"name"`
		MinCount int    `json:"min_count"`
	} `json:"required_tag_groups,omitempty"`
	ReadOnlyBanner                      string            `json:"read_only_banner"`
	AvailableGroups                     []string          `json:"available_groups,omitempty"`
	AutoCloseHours                      string            `json:"auto_close_hours"`
	AutoCloseBasedOnLastPost            bool              `json:"auto_close_based_on_last_post"`
	AllowUnlimitedOwnerEditsOnFirstPost bool              `json:"allow_unlimited_owner_edits_on_first_post"`
	DefaultSlowModeSeconds              string            `json:"default_slow_mode_seconds"`
	GroupPermissions                    []GroupPermission `json:"group_permissions"`
	EmailIn                             string            `json:"email_in"`
	EmailInAllowStrangers               bool              `json:"email_in_allow_strangers"`
	MailinglistMirror                   bool              `json:"mailinglist_mirror"`
	AllTopicsWiki                       bool              `json:"all_topics_wiki"`
	CanDelete                           bool              `json:"can_delete"`
	AllowBadges                         bool              `json:"allow_badges"`
	TopicFeaturedLinkAllowed            bool              `json:"topic_featured_link_allowed"`
	SearchPriority                      int               `json:"search_priority"`
	UploadedLogo                        string            `json:"uploaded_logo"`
	UploadedLogoDark                    string            `json:"uploaded_logo_dark"`
	UploadedBackground                  string            `json:"uploaded_background"`
	UploadedBackgroundDark              string            `json:"uploaded_background_dark"`
}

type ShowCategoryResponse struct {
	Category Category `json:"category"`
}

type ListCategoriesResponse struct {
	CategoryList CategoryList `json:"category_list"`
}

type CategoryList struct {
	CanCreateCategory bool       `json:"can_create_category"`
	CanCreateTopic    bool       `json:"can_create_topic"`
	Categories        []Category `json:"categories"`
}

func ShowCategory(client *Client, id int) (response *ShowCategoryResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("c/%d/show", id))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func ListCategories(client *Client) (response *ListCategoriesResponse, err error) {
	data, sendErr := client.Get("categories")

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func CreateCategory(client *Client, category *NewCategory) (response *ShowCategoryResponse, err error) {
	inputData, marshalError := json.Marshal(category)

	if marshalError != nil {
		return nil, marshalError
	}

	data, sendErr := client.PostWithReturn("categories", inputData)

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
