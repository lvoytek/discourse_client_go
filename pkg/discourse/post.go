package discourse

import (
	"encoding/json"
	"fmt"
)

type NewPost struct {
	Title             string `json:"title"`
	Raw               string `json:"raw"`
	TopicID           int    `json:"topic_id"`
	Category          int    `json:"category"`
	TargetRecipients  string `json:"target_recipients"`
	TargetUsernames   string `json:"target_usernames"`
	Archetype         string `json:"archetype"`
	CreatedAt         string `json:"created_at"`
	ReplyToPostNumber int    `json:"reply_to_post_number"`
	EmbedURL          string `json:"embed_url"`
	ExternalID        string `json:"external_id"`
}

type PostData struct {
	ID                          int          `json:"id"`
	Username                    string       `json:"username"`
	AvatarTemplate              string       `json:"avatar_template"`
	CreatedAt                   string       `json:"created_at"`
	Cooked                      string       `json:"cooked"`
	PostNumber                  int          `json:"post_number"`
	PostType                    int          `json:"post_type"`
	UpdatedAt                   string       `json:"updated_at"`
	ReplyCount                  int          `json:"reply_count"`
	ReplyToPostNumber           int          `json:"reply_to_post_number"`
	QuoteCount                  int          `json:"quote_count"`
	IncomingLinkCount           int          `json:"incoming_link_count"`
	Reads                       int          `json:"reads"`
	ReadersCount                int          `json:"readers_count"`
	Score                       float32      `json:"score"`
	Yours                       bool         `json:"yours"`
	TopicID                     int          `json:"topic_id"`
	TopicSlug                   string       `json:"topic_slug"`
	PrimaryGroupName            string       `json:"primary_group_name"`
	FlairName                   string       `json:"flair_name"`
	FlairURL                    string       `json:"flair_url"`
	FlairBgColor                string       `json:"flair_bg_color"`
	FlairColor                  string       `json:"flair_color"`
	FlairGroupID                int          `json:"flair_group_id"`
	Version                     int          `json:"version"`
	CanEdit                     bool         `json:"can_edit"`
	CanDelete                   bool         `json:"can_delete"`
	CanRecover                  bool         `json:"can_recover"`
	CanSeeHiddenPost            bool         `json:"can_see_hidden_post"`
	CanWiki                     bool         `json:"can_wiki"`
	UserTitle                   string       `json:"user_title"`
	Bookmarked                  bool         `json:"bookmarked"`
	Raw                         string       `json:"raw"`
	ActionsSummary              []PostAction `json:"actions_summary"`
	Moderator                   bool         `json:"moderator"`
	Admin                       bool         `json:"admin"`
	Staff                       bool         `json:"staff"`
	UserID                      int          `json:"user_id"`
	Hidden                      bool         `json:"hidden"`
	TrustLevel                  int          `json:"trust_level"`
	DeletedAt                   string       `json:"deleted_at"`
	UserDeleted                 bool         `json:"user_deleted"`
	EditReason                  string       `json:"edit_reason"`
	CanViewEditHistory          bool         `json:"can_view_edit_history"`
	Wiki                        bool         `json:"wiki"`
	ReviewableID                string       `json:"reviewable_id"`
	ReviewableScoreCount        int          `json:"reviewable_score_count"`
	ReviewableScorePendingCount int          `json:"reviewable_score_pending_count"`
	MentionedUsers              []string     `json:"mentioned_users"`
	Name                        string       `json:"name"`
	DisplayUsername             string       `json:"display_username"`
}

type PostAction struct {
	ID      int  `json:"id"`
	Count   int  `json:"count"`
	Acted   bool `json:"acted"`
	CanUndo bool `json:"can_undo"`
	CanAct  bool `json:"can_act"`
}

type GetLatestPostsResponse struct {
	LatestPosts []PostData `json:"latest_posts"`
}

func GetLatestPosts(client *Client) (response *GetLatestPostsResponse, err error) {
	data, sendErr := client.Get("posts", []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetPostByID(client *Client, id int) (response *PostData, err error) {
	data, sendErr := client.Get(fmt.Sprintf("posts/%d", id), []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetPostRepliesByID(client *Client, id int) (response []PostData, err error) {
	data, sendErr := client.Get(fmt.Sprintf("posts/%d/replies", id), []byte{})

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
