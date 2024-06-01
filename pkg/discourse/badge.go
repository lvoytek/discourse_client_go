package discourse

import (
	"encoding/json"
	"fmt"
)

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

type GrantedBy struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	AvatarTemplate string `json:"avatar_template"`
	FlairName      string `json:"flair_name"`
	Admin          bool   `json:"admin"`
	Moderator      bool   `json:"moderator"`
	TrustLevel     int    `json:"trust_level"`
}

type UserBadge struct {
	ID               int    `json:"id"`
	GrantedAt        string `json:"granted_at"`
	GroupingPosition int    `json:"grouping_position"`
	IsFavorite       string `json:"is_favorite"`
	CanFavorite      bool   `json:"can_favorite"`
	BadgeID          int    `json:"badge_id"`
	GrantedByID      int    `json:"granted_by_id"`
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

type ListBadgesForUserResponse struct {
	Badges      []Badge     `json:"badges"`
	BadgeTypes  []BadgeType `json:"badge_types"`
	GrantedBies []GrantedBy `json:"granted_bies"`
	UserBadges  []UserBadge `json:"user_badges"`
}

func ListBadgesForUser(client *Client, username string) (response *ListBadgesForUserResponse, err error) {
	data, sendErr := client.Get(fmt.Sprintf("user-badges/%s", username), []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
