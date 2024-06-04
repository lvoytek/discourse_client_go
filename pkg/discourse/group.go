package discourse

type Group struct {
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

type GetGroupResponse struct {
	ID                        int    `json:"id"`
	Automatic                 bool   `json:"automatic"`
	Name                      string `json:"name"`
	UserCount                 int    `json:"user_count"`
	MentionableLevel          int    `json:"mentionable_level"`
	MessageableLevel          int    `json:"messageable_level"`
	VisibilityLevel           int    `json:"visibility_level"`
	PrimaryGroup              bool   `json:"primary_group"`
	Title                     string `json:"title"`
	GrantTrustLevel           string `json:"grant_trust_level"`
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
