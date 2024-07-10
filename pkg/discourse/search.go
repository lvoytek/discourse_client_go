package discourse

import "time"

const (
	OrderLatest      string = "latest"
	OrderLikes              = "likes"
	OrderViews              = "views"
	OrderLatestTopic        = "latest_topic"
)

const (
	InTitle      string = "title"
	InLikes             = "likes"
	InPersonal          = "personal"
	InMessages          = "messages"
	InSeen              = "seen"
	InUnseen            = "unseen"
	InPosted            = "posted"
	InCreated           = "created"
	InWatching          = "watching"
	InTracking          = "tracking"
	InBookmarks         = "bookmarks"
	InAssigned          = "assigned"
	InUnassigned        = "unassigned"
	InFirst             = "first"
	InPinned            = "pinned"
	InWiki              = "wiki"
)

const (
	WithImages string = "images"
)

const (
	StatusOpen       string = "open"
	StatusClosed            = "closed"
	StatusPublic            = "public"
	StatusArchived          = "archived"
	StatusNoReplies         = "noreplies"
	StatusSingleUser        = "single_user"
	StatusSolved            = "solved"
	StatusUnsolved          = "unsolved"
)

type SearchQuery struct {
	Username         string
	Category         string
	Tags             []string
	Before           time.Time
	After            time.Time
	Order            string
	AssignedUsername string
	In               []string
	With             []string
	Status           []string
	Group            string
	GroupMessages    string
	MinPosts         int
	MaxPosts         int
	MinViews         int
	MaxViews         int
	Custom           map[string][]string
}
