package discourse

import "encoding/json"

type SiteInfo struct {
	DefaultArchetype                    string                   `json:"default_archetype"`
	NotificationTypes                   map[string]int           `json:"notification_types"`
	PostTypes                           map[string]int           `json:"post_types"`
	TrustLevels                         map[string]int           `json:"trust_levels"`
	UserTips                            map[string]int           `json:"user_tips"`
	Groups                              []Group                  `json:"groups"`
	Filters                             []string                 `json:"filters"`
	Periods                             []string                 `json:"periods"`
	TopMenuItems                        []string                 `json:"top_menu_items"`
	AnonymousTopMenuItems               []string                 `json:"anonymous_top_menu_items"`
	UncategorizedCategoryID             int                      `json:"uncategorized_category_id"`
	UserFieldMaxLength                  int                      `json:"user_field_max_length"`
	PostActionTypes                     []ActionType             `json:"post_action_types"`
	TopicFlagTypes                      []ActionType             `json:"topic_flag_types"`
	CanCreateTag                        bool                     `json:"can_create_tag"`
	CanTagTopics                        bool                     `json:"can_tag_topics"`
	CanTagPms                           bool                     `json:"can_tag_pms"`
	TagsFilterRegexp                    string                   `json:"tags_filter_regexp"`
	TopTags                             []string                 `json:"top_tags"`
	WizardRequired                      bool                     `json:"wizard_required"`
	CanAssociateGroups                  bool                     `json:"can_associate_groups"`
	TopicFeaturedLinkAllowedCategoryIDs []int                    `json:"topic_featured_link_allowed_category_ids"`
	UserThemes                          []SiteUserTheme          `json:"user_themes"`
	UserColorSchemes                    []SiteUserColorScheme    `json:"user_color_schemes"`
	DefaultDarkColorScheme              SiteUserColorScheme      `json:"default_dark_color_scheme"`
	CensoredRegexp                      []string                 `json:"censored_regexp"`
	CustomEmojiTranslation              map[string]interface{}   `json:"custom_emoji_translation"`
	WatchedWordsReplace                 string                   `json:"watched_words_replace"`
	WatchedWordsLink                    string                   `json:"watched_words_link"`
	MarkdownAdditionalOptions           map[string]interface{}   `json:"markdown_additional_options"`
	HashtagConfigurations               map[string]interface{}   `json:"hashtag_configurations"`
	HashtagIcons                        map[string]interface{}   `json:"hashtag_icons"`
	DisplayedAboutPluginStatGroups      []string                 `json:"displayed_about_plugin_stat_groups"`
	Categories                          []ShowCategoryResponse   `json:"categories"`
	Archetypes                          []SiteArchetype          `json:"archetypes"`
	UserFields                          []map[string]interface{} `json:"user_fields"`
	AuthProviders                       []map[string]interface{} `json:"auth_providers"`
	WhispersAllowedGroupsNames          []string                 `json:"whispers_allowed_groups_names"`
	DeniedEmojis                        []string                 `json:"denied_emojis"`
	NavigationMenuSiteTopTags           []string                 `json:"navigation_menu_site_top_tags"`
}

type SiteBasicInfo struct {
	LogoURL                    string `json:"logo_url"`
	LogoSmallURL               string `json:"logo_small_url"`
	AppleTouchIconURL          string `json:"apple_touch_icon_url"`
	FaviconURL                 string `json:"favicon_url"`
	Title                      string `json:"title"`
	Description                string `json:"description"`
	HeaderPrimaryColor         string `json:"header_primary_color"`
	HeaderBackgroundColor      string `json:"header_background_color"`
	LoginRequired              bool   `json:"login_required"`
	Locale                     string `json:"locale"`
	IncludeInDiscourseDiscover bool   `json:"include_in_discourse_discover"`
	MobileLogoURL              string `json:"mobile_logo_url"`
}

type ActionType struct {
	ID               int    `json:"id"`
	NameKey          string `json:"name_key"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	IsFlag           bool   `json:"is_flag"`
	IsCustomFlag     bool   `json:"is_custom_flag"`
}

type SiteUserTheme struct {
	ThemeID       int    `json:"theme_id"`
	Name          string `json:"name"`
	Default       bool   `json:"default"`
	ColorSchemeID int    `json:"color_scheme_id"`
}

type SiteUserColorScheme struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDark bool   `json:"is_dark"`
}

type SiteArchetype struct {
	ID      string              `json:"id"`
	Name    string              `json:"name"`
	Options []map[string]string `json:"options"`
}

func GetSiteBasicInfo(client *Client) (response *SiteBasicInfo, err error) {
	data, sendErr := client.Get("site/basic-info")

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetSiteInfo(client *Client) (response *SiteInfo, err error) {
	data, sendErr := client.Get("site")

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
