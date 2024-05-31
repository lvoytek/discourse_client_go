package discourse

type Badge struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	GrantCount        int    `json:"grant_count"`
	AllowTitle        bool   `json:"allow_title"`
	MultipleGrant     bool   `json:"multiple_grant"`
	Icon              string `json:"icon"`
	ImageURL          string `json:"image_url"`
	Listable          bool   `json:"listable"`
	Enabled           bool   `json:"enabled"`
	BadgeGroupingID   int    `json:"badge_grouping_id"`
	System            bool   `json:"system"`
	LongDescription   string `json:"long_description"`
	Slug              string `json:"slug"`
	ManuallyGrantable bool   `json:"manually_grantable"`
	Query             string `json:"query"`
	Trigger           string `json:"trigger"`
	TargetPosts       bool   `json:"target_posts"`
	AutoRevoke        bool   `json:"auto_revoke"`
	ShowPosts         bool   `json:"show_posts"`
	BadgeTypeID       int    `json:"badge_type_id"`
}

type BadgeType struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type BadgeGrouping struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	System      bool   `json:"system"`
}

type ListBadgesResponse struct {
	Badges         []Badge         `json:"badges"`
	BadgeTypes     []BadgeType     `json:"badge_types"`
	BadgeGroupings []BadgeGrouping `json:"badge_groupings"`
	AdminBadges    []struct {
		ProtectedSystemFields []string `json:"protected_system_fields,omitempty"`
		Triggers              []struct {
			UserChange       int `json:"user_change"`
			None             int `json:"none"`
			PostRevision     int `json:"post_revision"`
			TrustLevelChange int `json:"trust_level_change"`
			PostAction       int `json:"post_action"`
		} `json:"triggers,omitempty"`
		BadgeIDs         []int `json:"badge_ids,omitempty"`
		BadgeGroupingIDs []int `json:"badge_grouping_ids,omitempty"`
		BadgeTypeIDs     []int `json:"badge_type_ids,omitempty"`
	} `json:"admin_badges,omitempty"`
}
