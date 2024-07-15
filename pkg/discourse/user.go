package discourse

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type NewUser struct {
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Username    string            `json:"username"`
	Active      bool              `json:"active,omitempty"`
	Approved    bool              `json:"approved,omitempty"`
	UserFields  map[string]bool   `json:"user_fields,omitempty"`
	ExternalIDs map[string]string `json:"external_ids,omitempty"`
}

type User struct {
	ID                             int                    `json:"id"`
	Username                       string                 `json:"username"`
	Name                           string                 `json:"name"`
	LastPostedAt                   string                 `json:"last_posted_at"`
	LastSeenAt                     string                 `json:"last_seen_at"`
	CreatedAt                      string                 `json:"created_at"`
	Ignored                        bool                   `json:"ignored"`
	Muted                          bool                   `json:"muted"`
	CanIgnoreUser                  bool                   `json:"can_ignore_user"`
	CanMuteUser                    bool                   `json:"can_mute_user"`
	CanSendPrivateMessages         bool                   `json:"can_send_private_messages"`
	CanSendPrivateMessageToUser    bool                   `json:"can_send_private_message_to_user"`
	TrustLevel                     int                    `json:"trust_level"`
	Moderator                      bool                   `json:"moderator"`
	Admin                          bool                   `json:"admin"`
	Title                          string                 `json:"title"`
	BadgeCount                     int                    `json:"badge_count"`
	SecondFactorBackupEnabled      bool                   `json:"second_factor_backup_enabled"`
	UserFields                     map[string]interface{} `json:"user_fields"`
	CustomFields                   map[string]interface{} `json:"custom_fields"`
	TimeRead                       int                    `json:"time_read"`
	RecentTimeRead                 int                    `json:"recent_time_read"`
	PrimaryGroupID                 string                 `json:"primary_group_id"`
	PrimaryGroupName               string                 `json:"primary_group_name"`
	FlairGroupID                   string                 `json:"flair_group_id"`
	FlairName                      string                 `json:"flair_name"`
	FlairURL                       string                 `json:"flair_url"`
	FlairBgColor                   string                 `json:"flair_bg_color"`
	FlairColor                     string                 `json:"flair_color"`
	FeaturedTopic                  string                 `json:"featured_topic"`
	Staged                         bool                   `json:"staged"`
	CanEdit                        bool                   `json:"can_edit"`
	CanEditUsername                bool                   `json:"can_edit_username"`
	CanEditEmail                   bool                   `json:"can_edit_email"`
	CanEditName                    bool                   `json:"can_edit_name"`
	UploadedAvatarID               string                 `json:"uploaded_avatar_id"`
	HasTitleBadges                 bool                   `json:"has_title_badges"`
	PendingCount                   int                    `json:"pending_count"`
	PendingPostsCount              int                    `json:"pending_posts_count"`
	ProfileViewCount               int                    `json:"profile_view_count"`
	SecondFactorEnabled            bool                   `json:"second_factor_enabled"`
	CanUploadProfileHeader         bool                   `json:"can_upload_profile_header"`
	CanUploadUserCardBackground    bool                   `json:"can_upload_user_card_background"`
	PostCount                      int                    `json:"post_count"`
	CanBeDeleted                   bool                   `json:"can_be_deleted"`
	CanDeleteAllPosts              bool                   `json:"can_delete_all_posts"`
	Locale                         string                 `json:"locale"`
	MutedCategoryIDs               []int                  `json:"muted_category_ids"`
	RegularCategoryIDs             []int                  `json:"regular_category_ids"`
	WatchedTags                    []int                  `json:"watched_tags"`
	WatchingFirstPostTags          []int                  `json:"watching_first_post_tags"`
	TrackedTags                    []int                  `json:"tracked_tags"`
	MutedTags                      []int                  `json:"muted_tags"`
	TrackedCategoryIDs             []int                  `json:"tracked_category_ids"`
	WatchedCategoryIDs             []int                  `json:"watched_category_ids"`
	WatchedFirstPostCategoryIDs    []int                  `json:"watched_first_post_category_ids"`
	SystemAvatarUploadID           string                 `json:"system_avatar_upload_id"`
	SystemAvatarTemplate           string                 `json:"system_avatar_template"`
	MutedUsernames                 []string               `json:"muted_usernames"`
	IgnoredUsernames               []string               `json:"ignored_usernames"`
	AllowedPmUsernames             []string               `json:"allowed_pm_usernames"`
	MailingListPostsPerDay         int                    `json:"mailing_list_posts_per_day"`
	CanChangeBio                   bool                   `json:"can_change_bio"`
	CanChangeLocation              bool                   `json:"can_change_location"`
	CanChangeWebsite               bool                   `json:"can_change_website"`
	CanChangeTrackingPreferences   bool                   `json:"can_change_tracking_preferences"`
	UserApiKeys                    string                 `json:"user_api_keys"`
	UserPasskeys                   []string               `json:"user_passkeys"`
	SidebarTags                    []string               `json:"sidebar_tags"`
	SidebarCategoryIDs             []int                  `json:"sidebar_category_ids"`
	DisplaySidebarTags             bool                   `json:"display_sidebar_tags"`
	CanPickThemeWithCustomHomepage bool                   `json:"can_pick_theme_with_custom_homepage"`
	UserAuthTokens                 []UserAuthToken        `json:"user_auth_tokens"`
	UserNotificationSchedule       map[string]interface{} `json:"user_notification_schedule"`
	UseLogoSmallAsAvatar           bool                   `json:"use_logo_small_as_avatar"`
	FeaturedUserBadgeIDs           []int                  `json:"featured_user_badge_ids"`
	InvitedBy                      string                 `json:"invited_by"`
	Groups                         []Group                `json:"groups"`
	GroupUsers                     []GroupUser            `json:"group_users"`
	UserOption                     UserOption             `json:"user_option"`
}

type UserAuthToken struct {
	ID        int    `json:"id"`
	ClientIp  string `json:"client_ip"`
	Location  string `json:"location"`
	Browser   string `json:"browser"`
	Device    string `json:"device"`
	OS        string `json:"os"`
	Icon      string `json:"icon"`
	CreatedAt string `json:"created_at"`
	SeenAt    string `json:"seen_at"`
	IsActive  bool   `json:"is_active"`
}

type UserOption struct {
	UserID                        int    `json:"user_id"`
	MailingListMode               bool   `json:"mailing_list_mode"`
	MailingListModeFrequency      int    `json:"mailing_list_mode_frequency"`
	EmailDigests                  bool   `json:"email_digests"`
	EmailLevel                    int    `json:"email_level"`
	EmailMessagesLevel            int    `json:"email_messages_level"`
	ExternalLinksInNewTab         bool   `json:"external_links_in_new_tab"`
	BookmarkAutoDeletePreference  int    `json:"bookmark_auto_delete_preference"`
	ColorSchemeID                 string `json:"color_scheme_id"`
	DarkSchemeID                  string `json:"dark_scheme_id"`
	DynamicFavicon                bool   `json:"dynamic_favicon"`
	EnableQuoting                 bool   `json:"enable_quoting"`
	EnableDefer                   bool   `json:"enable_defer"`
	DigestAfterMinutes            int    `json:"digest_after_minutes"`
	AutomaticallyUnpinTopics      bool   `json:"automatically_unpin_topics"`
	AutoTrackTopicsAfterMsecs     int    `json:"auto_track_topics_after_msecs"`
	NotificationLevelWhenReplying int    `json:"notification_level_when_replying"`
	NewTopicDurationMinutes       int    `json:"new_topic_duration_minutes"`
	EmailPreviousReplies          int    `json:"email_previous_replies"`
	EmailInReplyTo                bool   `json:"email_in_reply_to"`
	LikeNotificationFrequency     int    `json:"like_notification_frequency"`
	IncludeTl0InDigests           bool   `json:"include_tl0_in_digests"`
	ThemeIDs                      []int  `json:"theme_ids"`
	ThemeKeySeq                   int    `json:"theme_key_seq"`
	AllowPrivateMessages          bool   `json:"allow_private_messages"`
	EnableAllowedPmUsers          bool   `json:"enable_allowed_pm_users"`
	HomepageID                    string `json:"homepage_id"`
	HideProfileAndPresence        bool   `json:"hide_profile_and_presence"`
	TextSize                      string `json:"text_size"`
	TextSizeSeq                   int    `json:"text_size_seq"`
	TitleCountMode                string `json:"title_count_mode"`
	Timezone                      string `json:"timezone"`
	SkipNewUserTips               bool   `json:"skip_new_user_tips"`
	DefaultCalendar               string `json:"default_calendar"`
	OldestSearchLogDate           string `json:"oldest_search_log_date"`
	SidebarLinkToFilteredList     bool   `json:"sidebar_link_to_filtered_list"`
	SidebarShowCountOfNewItems    bool   `json:"sidebar_show_count_of_new_items"`
	WatchedPrecedenceOverMuted    bool   `json:"watched_precedence_over_muted"`
	SeenPopups                    []int  `json:"seen_popups"`
	TopicsUnreadWhenClosed        bool   `json:"topics_unread_when_closed"`
}

type CreateUserResponse struct {
	Success bool   `json:"success"`
	Active  bool   `json:"active"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

type UpdateUserResponse struct {
	Success string `json:"success"`
	User    User   `json:"user"`
}

type GetUserResponse struct {
	UserBadges []UserBadge `json:"user_badges"`
	User       User        `json:"user"`
}

const (
	UserAvatarChoiceUploaded string = "uploaded"
	UserAvatarChoiceCustom   string = "custom"
	UserAvatarChoiceGravatar string = "gravatar"
	UserAvatarChoiceSystem   string = "system"
)

type UserAvatarChoice struct {
	UploadID int    `json:"upload_id"`
	Type     string `json:"type"`
}

type UserEmail struct {
	Email string `json:"email"`
}

type UserNewUsername struct {
	Username string `json:"new_username"`
}

type UserDeleteOptions struct {
	DeletePosts bool `json:"delete_posts"`
	BlockEmail  bool `json:"block_email"`
	BlockURLs   bool `json:"block_urls"`
	BlockIP     bool `json:"block_ip"`
}

const (
	PeriodDaily     string = "daily"
	PeriodWeekly    string = "weekly"
	PeriodMonthly   string = "monthly"
	PeriodQuarterly string = "quarterly"
	PeriodYearly    string = "yearly"
	PeriodAll       string = "all"
)

const (
	OrderLikesReceived string = "likes_received"
	OrderLikesGiven    string = "likes_given"
	OrderTopicCount    string = "topic_count"
	OrderPostCount     string = "post_count"
	OrderTopicsEntered string = "topics_entered"
	OrderPostsRead     string = "posts_read"
	OrderDaysVisited   string = "days_visited"
)

type UserDirectoryQuery struct {
	Period    string
	Order     string
	Ascending bool
	Page      int
}

type UserDirectory struct {
	DirectoryItems []UserDirectoryEntry  `json:"directory_items"`
	Meta           UserDirectoryMetadata `json:"meta"`
}

type UserDirectoryEntry struct {
	ID            int  `json:"id"`
	LikesReceived int  `json:"likes_received"`
	LikesGiven    int  `json:"likes_given"`
	TopicsEntered int  `json:"topics_entered"`
	TopicCount    int  `json:"topic_count"`
	PostCount     int  `json:"post_count"`
	PostsRead     int  `json:"posts_read"`
	DaysVisited   int  `json:"days_visited"`
	User          User `json:"user"`
}

type UserDirectoryMetadata struct {
	LastUpdatedAt           string `json:"last_updated_at"`
	TotalRowsDirectoryItems int    `json:"total_rows_directory_items"`
	LoadMoreDirectoryItems  string `json:"load_more_directory_items"`
}

func CreateUser(client *Client, user *NewUser) (response *CreateUserResponse, err error) {
	inputData, marshalError := json.Marshal(user)

	if marshalError != nil {
		return nil, marshalError
	}

	data, sendErr := client.PostWithReturn("users", inputData)

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetUserByUsername(client *Client, username string) (response *GetUserResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("u/%s", username))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetUserByExternalID(client *Client, externalID string) (response *GetUserResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("u/by-external/%s", externalID))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetUserByExternalAuthID(client *Client, provider string, externalID string) (response *GetUserResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("u/by-external/%s/%s", provider, externalID))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func ListUsersByStats(client *Client, query *UserDirectoryQuery) (response *UserDirectory, err error) {
	data, sendErr := client.GetWithQueryString("directory_items", url.QueryEscape(generateListUsersQuery(query)))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func UpdateUserByUsername(client *Client, username string, userData *NewUser) (response *UpdateUserResponse, err error) {
	inputData, marshalError := json.Marshal(userData)

	if marshalError != nil {
		return nil, marshalError
	}

	data, sendErr := client.PutWithReturn(fmt.Sprintf("u/%s", username), inputData)

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func UpdateUserAvatarByUsername(client *Client, username string, avatarChoice *UserAvatarChoice) error {
	inputData, marshalError := json.Marshal(avatarChoice)

	if marshalError != nil {
		return marshalError
	}

	return client.Put(fmt.Sprintf("u/%s/preferences/avatar/pick", username), inputData)
}

func UpdateUserEmailByUsername(client *Client, username string, email string) error {
	inputData, marshalError := json.Marshal(UserEmail{Email: email})

	if marshalError != nil {
		return marshalError
	}

	return client.Put(fmt.Sprintf("u/%s/preferences/email", username), inputData)
}

func UpdateUserUsernameByUsername(client *Client, username string, newUsername string) error {
	inputData, marshalError := json.Marshal(UserNewUsername{Username: newUsername})

	if marshalError != nil {
		return marshalError
	}

	return client.Put(fmt.Sprintf("u/%s/preferences/username", username), inputData)
}

func DeleteUserByID(client *Client, id int, deleteOptions *UserDeleteOptions) error {
	inputData, marshalError := json.Marshal(deleteOptions)

	if marshalError != nil {
		return marshalError
	}

	return client.Delete(fmt.Sprintf("admin/users/%d", id), inputData)
}

func ActivateUserByID(client *Client, id int) error {
	return client.Put(fmt.Sprintf("admin/users/%d/activate", id), []byte{})
}

func DeactivateUserByID(client *Client, id int) error {
	return client.Put(fmt.Sprintf("admin/users/%d/deactivate", id), []byte{})
}

func generateListUsersQuery(query *UserDirectoryQuery) string {
	if query.Period == "" {
		query.Period = "all"
	}

	searchString := fmt.Sprintf("period=%s", query.Period)

	if query.Order != "" {
		searchString = fmt.Sprintf("%s&order=%s", searchString, query.Order)
	}

	if query.Ascending {
		searchString += "&asc=true"
	}

	if query.Page > 0 {
		searchString = fmt.Sprintf("%s&page=%d", searchString, query.Page)
	}

	return searchString
}
