package discourse

import (
	"encoding/json"
	"fmt"
)

type TopicData struct {
	PostStream           PostStream        `json:"post_stream"`
	TimelineLookup       [][]int           `json:"timeline_lookup"`
	SuggestedTopics      []SuggestedTopic  `json:"suggested_topics"`
	Tags                 []string          `json:"tags"`
	TagsDescriptions     map[string]string `json:"tags_descriptions"`
	ID                   int               `json:"id"`
	Title                string            `json:"title"`
	FancyTitle           string            `json:"fancy_title"`
	PostsCount           int               `json:"posts_count"`
	CreatedAt            string            `json:"created_at"`
	Views                int               `json:"views"`
	ReplyCount           int               `json:"reply_count"`
	LikeCount            int               `json:"like_count"`
	LastPostedAt         string            `json:"last_posted_at"`
	Visible              bool              `json:"visible"`
	Closed               bool              `json:"closed"`
	Archived             bool              `json:"archived"`
	HasSummary           bool              `json:"has_summary"`
	Archetype            string            `json:"archetype"`
	Slug                 string            `json:"slug"`
	CategoryID           int               `json:"category_id"`
	WordCount            int               `json:"word_count"`
	DeletedAt            string            `json:"deleted_at"`
	UserID               int               `json:"user_id"`
	FeaturedLink         string            `json:"featured_link"`
	PinnedGlobally       bool              `json:"pinned_globally"`
	PinnedAt             string            `json:"pinned_at"`
	PinnedUntil          string            `json:"pinned_until"`
	ImageURL             string            `json:"image_url"`
	SlowModeSeconds      int               `json:"slow_mode_seconds"`
	Draft                string            `json:"draft"`
	DraftKey             string            `json:"draft_key"`
	DraftSequence        int               `json:"draft_sequence"`
	Unpinned             string            `json:"unpinned"`
	Pinned               bool              `json:"pinned"`
	CurrentPostNumber    int               `json:"current_post_number"`
	HighestPostNumber    int               `json:"highest_post_number"`
	DeletedBy            string            `json:"deleted_by"`
	HasDeleted           bool              `json:"has_deleted"`
	ActionsSummary       []TopicAction     `json:"actions_summary"`
	ChunkSize            int               `json:"chunk_size"`
	Bookmarked           bool              `json:"bookmarked"`
	Bookmarks            []string          `json:"bookmarks"`
	TopicTimer           string            `json:"topic_timer"`
	MessageBusLastID     int               `json:"message_bus_last_id"`
	ParticipantCount     int               `json:"participant_count"`
	ShowReadIndicator    bool              `json:"show_read_indicator"`
	Thumbnails           []Thumbnail       `json:"thumbnails"`
	SlowModeEnabledUntil string            `json:"slow_mode_enabled_until"`
	Summarizable         bool              `json:"summarizable"`
	Details              TopicDetails      `json:"details"`
}

type SuggestedTopic struct {
	ID                int               `json:"id"`
	Title             string            `json:"title"`
	FancyTitle        string            `json:"fancy_title"`
	Slug              string            `json:"slug"`
	PostsCount        int               `json:"posts_count"`
	ReplyCount        int               `json:"reply_count"`
	HighestPostNumber int               `json:"highest_post_number"`
	ImageURL          string            `json:"image_url"`
	CreatedAt         string            `json:"created_at"`
	LastPostedAt      string            `json:"last_posted_at"`
	Bumped            bool              `json:"bumped"`
	BumpedAt          string            `json:"bumped_at"`
	Archetype         string            `json:"archetype"`
	Unseen            bool              `json:"unseen"`
	Pinned            bool              `json:"pinned"`
	Unpinned          string            `json:"unpinned"`
	Excerpt           string            `json:"excerpt"`
	Visible           bool              `json:"visible"`
	Closed            bool              `json:"closed"`
	Archived          bool              `json:"archived"`
	Bookmarked        string            `json:"bookmarked"`
	Liked             string            `json:"liked"`
	Tags              []string          `json:"tags"`
	TagsDescriptions  map[string]string `json:"tags_descriptions"`
	LikeCount         int               `json:"like_count"`
	Views             int               `json:"views"`
	CategoryID        int               `json:"category_id"`
	FeaturedLink      string            `json:"featured_link"`
	Posters           []TopicPoster     `json:"posters"`
}

type TopicDetails struct {
	CanEdit                  bool               `json:"can_edit"`
	NotificationLevel        int                `json:"notification_level"`
	CanMovePosts             bool               `json:"can_move_posts"`
	CanDelete                bool               `json:"can_delete"`
	CanRemoveAllowedUsers    bool               `json:"can_remove_allowed_users"`
	CanCreatePost            bool               `json:"can_create_post"`
	CanReplyAsNewTopic       bool               `json:"can_reply_as_new_topic"`
	CanInviteTo              bool               `json:"can_invite_to"`
	CanInviteViaEmail        bool               `json:"can_invite_via_email"`
	CanFlagTopic             bool               `json:"can_flag_topic"`
	CanConvertTopic          bool               `json:"can_convert_topic"`
	CanReviewTopic           bool               `json:"can_review_topic"`
	CanCloseTopic            bool               `json:"can_close_topic"`
	CanArchiveTopic          bool               `json:"can_archive_topic"`
	CanSplitMergeTopic       bool               `json:"can_split_merge_topic"`
	CanEditStaffNotes        bool               `json:"can_edit_staff_notes"`
	CanToggleTopicVisibility bool               `json:"can_toggle_topic_visibility"`
	CanPinUnpinTopic         bool               `json:"can_pin_unpin_topic"`
	CanModerateCategory      bool               `json:"can_moderate_category"`
	CanRemoveSelfID          int                `json:"can_remove_self_id"`
	Participants             []TopicParticipant `json:"participants"`
	CreatedBy                PostCreator        `json:"created_by"`
	LastPoster               PostCreator        `json:"last_poster"`
}

type PostStream struct {
	Posts  []PostData `json:"posts"`
	Stream []int      `json:"stream"`
}

type TopicPoster struct {
	Extras      string      `json:"extras"`
	Description string      `json:"description"`
	User        PostCreator `json:"user"`
}

type TopicParticipant struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	AvatarTemplate   string `json:"avatar_template"`
	PostCount        int    `json:"post_count"`
	PrimaryGroupName string `json:"primary_group_name"`
	FlairName        string `json:"flair_name"`
	FlairURL         string `json:"flair_url"`
	FlairColor       string `json:"flair_color"`
	FlairBgColor     string `json:"flair_bg_color"`
	FlairGroupID     string `json:"flair_group_id"`
	Admin            bool   `json:"admin"`
	Moderator        bool   `json:"moderator"`
	TrustLevel       int    `json:"trust_level"`
}

type TopicAction struct {
	ID     int  `json:"id"`
	Count  int  `json:"count"`
	Hidden bool `json:"hidden"`
	CanAct bool `json:"can_act"`
}

type Thumbnail struct {
	MaxWidth  int    `json:"max_width"`
	MaxHeight int    `json:"max_height"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	URL       string `json:"url"`
}

func GetTopicByID(client *Client, id int) (response *TopicData, err error) {
	data, sendErr := client.Get(fmt.Sprintf("t/%d", id), []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
