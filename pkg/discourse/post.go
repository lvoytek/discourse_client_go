package discourse

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type NewPost struct {
	Title             string    `json:"title"`
	Raw               string    `json:"raw"`
	TopicID           int       `json:"topic_id,omitempty"`
	Category          int       `json:"category,omitempty"`
	TargetRecipients  string    `json:"target_recipients,omitempty"`
	TargetUsernames   string    `json:"target_usernames,omitempty"`
	Archetype         string    `json:"archetype"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	ReplyToPostNumber int       `json:"reply_to_post_number,omitempty"`
	EmbedURL          string    `json:"embed_url,omitempty"`
	ExternalID        string    `json:"external_id,omitempty"`
}

type PostData struct {
	ID                          int          `json:"id"`
	Username                    string       `json:"username"`
	AvatarTemplate              string       `json:"avatar_template"`
	CreatedAt                   time.Time    `json:"created_at"`
	Cooked                      string       `json:"cooked"`
	PostNumber                  int          `json:"post_number"`
	PostType                    int          `json:"post_type"`
	UpdatedAt                   time.Time    `json:"updated_at"`
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
	DeletedAt                   time.Time    `json:"deleted_at"`
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

type PostRevision struct {
	CreatedAt        time.Time       `json:"created_at"`
	PostID           int             `json:"post_id"`
	PreviousHidden   bool            `json:"previous_hidden"`
	CurrentHidden    bool            `json:"current_hidden"`
	FirstRevision    int             `json:"first_revision"`
	PreviousRevision int             `json:"previous_revision"`
	CurrentRevision  int             `json:"current_revision"`
	NextRevision     int             `json:"next_revision"`
	LastRevision     int             `json:"last_revision"`
	CurrentVersion   int             `json:"current_version"`
	VersionCount     int             `json:"version_count"`
	Username         string          `json:"username"`
	DisplayUsername  string          `json:"display_username"`
	AvatarTemplate   string          `json:"avatar_template"`
	EditReason       string          `json:"edit_reason"`
	BodyChanges      RevisionChanges `json:"body_changes,omitempty"`
	TitleChanges     RevisionChanges `json:"title_changes,omitempty"`
	Wiki             bool            `json:"wiki"`
	CanEdit          bool            `json:"can_edit"`
}

type RevisionChanges struct {
	Inline             string `json:"inline"`
	SideBySide         string `json:"side_by_side"`
	SideBySideMarkdown string `json:"side_by_side_markdown,omitempty"`
}

type PostAction struct {
	ID      int  `json:"id"`
	Count   int  `json:"count"`
	Acted   bool `json:"acted"`
	CanUndo bool `json:"can_undo"`
	CanAct  bool `json:"can_act"`
}

type PostCreator struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	AvatarTemplate string `json:"avatar_template"`
	Name           string `json:"name"`
}

type GetLatestPostsResponse struct {
	LatestPosts []PostData `json:"latest_posts"`
}

func GetLatestPosts(client *Client) (response *GetLatestPostsResponse, err error) {
	data, sendErr := client.Get("posts")

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetPostByID(client *Client, id int) (response *PostData, err error) {
	data, sendErr := client.Get(fmt.Sprintf("posts/%d", id))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetPostRepliesByID(client *Client, id int) (response []PostData, err error) {
	data, sendErr := client.Get(fmt.Sprintf("posts/%d/replies", id))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetPostRevisionByID(client *Client, id int, revisionNumber int) (response *PostRevision, err error) {
	data, sendErr := client.Get(fmt.Sprintf("posts/%d/revisions/%d", id, revisionNumber))

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

func GetNumPostRevisionsByID(client *Client, id int) (response int, err error) {
	secondRevision, err := GetPostRevisionByID(client, id, 2)

	if err == nil {
		return secondRevision.VersionCount, nil
	} else if strings.Contains(fmt.Sprint(err), "404") || strings.Contains(fmt.Sprint(err), "403") {
		_, postExistsErr := GetPostByID(client, id)

		if postExistsErr == nil {
			return 1, nil
		}
	}

	return 0, err
}

func CreatePost(client *Client, post *NewPost) (response *PostData, err error) {
	post.Archetype = "post"
	inputData, marshalError := json.Marshal(post)

	if marshalError != nil {
		return nil, marshalError
	}

	data, sendErr := client.PostWithReturn("posts", inputData)

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
