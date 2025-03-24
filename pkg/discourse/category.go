package discourse

import (
	"encoding/json"
	"fmt"
)

type NewCategory struct {
	Name                      string         `json:"name,omitempty"`
	Color                     string         `json:"color,omitempty"`
	TextColor                 string         `json:"text_color,omitempty"`
	ParentCategoryID          int            `json:"parent_category_id,omitempty"`
	AllowBadges               bool           `json:"allow_badges,omitempty"`
	Slug                      string         `json:"slug,omitempty"`
	TopicFeaturedLinksAllowed bool           `json:"topic_featured_links_allowed,omitempty"`
	Permissions               map[string]int `json:"permissions,omitempty"`
	SearchPriority            int            `json:"search_priority,omitempty"`
	FormTemplateIDs           map[string]int `json:"form_template_ids,omitempty"`
}

type CategoryWithoutSubcategories struct {
	ID                           int                    `json:"id"`
	Name                         string                 `json:"name"`
	Color                        string                 `json:"color"`
	TextColor                    string                 `json:"text_color"`
	Slug                         string                 `json:"slug"`
	TopicCount                   int                    `json:"topic_count"`
	PostCount                    int                    `json:"post_count"`
	Position                     int                    `json:"position"`
	Description                  string                 `json:"description"`
	DescriptionText              string                 `json:"description_text"`
	DescriptionExcerpt           string                 `json:"description_excerpt"`
	TopicURL                     string                 `json:"topic_url"`
	ReadRestricted               bool                   `json:"read_restricted"`
	Permission                   int                    `json:"permission"`
	NotificationLevel            int                    `json:"notification_level"`
	CanEdit                      bool                   `json:"can_edit"`
	TopicTemplate                string                 `json:"topic_template"`
	FormTemplateIDs              []int                  `json:"form_template_ids"`
	HasChildren                  bool                   `json:"has_children"`
	SubcategoryCount             int                    `json:"subcategory_count"`
	SortOrder                    string                 `json:"sort_order"`
	SortAscending                string                 `json:"sort_ascending"`
	ShowSubcategoryList          bool                   `json:"show_subcategory_list"`
	NumFeaturedTopics            int                    `json:"num_featured_topics"`
	DefaultView                  string                 `json:"default_view"`
	SubcategoryListStyle         string                 `json:"subcategory_list_style"`
	DefaultTopPeriod             string                 `json:"default_top_period"`
	DefaultListFilter            string                 `json:"default_list_filter"`
	MinimumRequiredTags          int                    `json:"minimum_required_tags"`
	NavigateToFirstPostAfterRead bool                   `json:"navigate_to_first_post_after_read"`
	UploadedLogo                 Image                  `json:"uploaded_logo"`
	UploadedLogoDark             Image                  `json:"uploaded_logo_dark"`
	UploadedBackground           Image                  `json:"uploaded_background"`
	UploadedBackgroundDark       Image                  `json:"uploaded_background_dark"`
	CustomFields                 map[string]interface{} `json:"custom_fields,omitempty"`
	AllowedTags                  []Tag                  `json:"allowed_tags,omitempty"`
	AllowedTagGroups             []string               `json:"allowed_tag_groups,omitempty"`
	AllowGlobalTags              bool                   `json:"allow_global_tags,omitempty"`
	RequiredTagGroups            []struct {
		Name     string `json:"name"`
		MinCount int    `json:"min_count"`
	} `json:"required_tag_groups,omitempty"`
	ReadOnlyBanner                      string            `json:"read_only_banner,omitempty"`
	AvailableGroups                     []string          `json:"available_groups,omitempty"`
	AutoCloseHours                      string            `json:"auto_close_hours,omitempty"`
	AutoCloseBasedOnLastPost            bool              `json:"auto_close_based_on_last_post,omitempty"`
	AllowUnlimitedOwnerEditsOnFirstPost bool              `json:"allow_unlimited_owner_edits_on_first_post,omitempty"`
	DefaultSlowModeSeconds              string            `json:"default_slow_mode_seconds,omitempty"`
	GroupPermissions                    []GroupPermission `json:"group_permissions,omitempty"`
	EmailIn                             string            `json:"email_in,omitempty"`
	EmailInAllowStrangers               bool              `json:"email_in_allow_strangers,omitempty"`
	MailinglistMirror                   bool              `json:"mailinglist_mirror,omitempty"`
	AllTopicsWiki                       bool              `json:"all_topics_wiki,omitempty"`
	CanDelete                           bool              `json:"can_delete,omitempty"`
	AllowBadges                         bool              `json:"allow_badges,omitempty"`
	TopicFeaturedLinkAllowed            bool              `json:"topic_featured_link_allowed,omitempty"`
	SearchPriority                      int               `json:"search_priority,omitempty"`
	TopicsDay                           int               `json:"topics_day,omitempty"`
	TopicsWeek                          int               `json:"topics_week,omitempty"`
	TopicsMonth                         int               `json:"topics_month,omitempty"`
	TopicsYear                          int               `json:"topics_year,omitempty"`
	TopicsAllTime                       int               `json:"topics_all_time,omitempty"`
	IsUncategorized                     bool              `json:"is_uncategorized,omitempty"`
	SubcategoryIDs                      []int             `json:"subcategory_ids,omitempty"`
}

type Category struct {
	CategoryWithoutSubcategories
	SubcategoryList []CategoryWithoutSubcategories `json:"subcategory_list"`
}

type CategoryContents struct {
	Users         []User            `json:"users"`
	PrimaryGroups []Group           `json:"primary_groups"`
	TopicList     CategoryTopicList `json:"topic_list"`
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

type CategoryTopicList struct {
	CanCreateTopic bool             `json:"can_create_topic"`
	PerPage        int              `json:"per_page"`
	TopTags        []string         `json:"top_tags"`
	Topics         []SuggestedTopic `json:"topics"`
}

func ShowCategory(client *Client, id int) (response *ShowCategoryResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("c/%d/show", id))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetCategoryContentsByID(client *Client, id int, page int) (response *CategoryContents, err error) {
	data, sendErr := client.GetWithQueryString(fmt.Sprintf("c/%d", id), fmt.Sprintf("page=%d", page))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetCategoryContentsBySlug(client *Client, slug string, page int) (response *CategoryContents, err error) {
	data, sendErr := client.GetWithQueryString(fmt.Sprintf("c/%s", slug), fmt.Sprintf("page=%d", page))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func ListCategories(client *Client, showSubcategories bool) (response *ListCategoriesResponse, err error) {
	data, sendErr := client.GetWithQueryString("categories", fmt.Sprintf("include_subcategories=%t", showSubcategories))

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

func UpdateCategoryByID(client *Client, id int, categoryData *NewCategory) (response *ShowCategoryResponse, err error) {
	inputData, marshalError := json.Marshal(categoryData)

	if marshalError != nil {
		return nil, marshalError
	}

	data, sendErr := client.PutWithReturn(fmt.Sprintf("categories/%d", id), inputData)

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
