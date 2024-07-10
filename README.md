# Discourse API Client for Go
Library for interacting with a Discourse site via the [Discourse API](https://docs.discourse.org/)

## Installation
Download the library to your project with `go get`:

```bash
go get github.com/lvoytek/discourse_client_go
```
## Usage
### Initialization
All interactions with a site using this library require a `Client` variable which can be created with either `NewClient` or `NewAnonymousClient`. An API key and username are needed to upload data or access admin data, which can be provided with `NewClient`:

```go
discourseClient := discourse.NewClient("http://127.0.0.1:3000/", "714552c...", "system")
```
Most downloading can be done without an API key. If you would just like to get data from the Discourse site, then you can create an anonymous client:

```go
discourseClient := discourse.NewAnonymousClient("https://discourse.ubuntu.com")
```
### Access

Functions that access the Discourse site are meant to match [Discourse API](https://docs.discourse.org/) calls 1:1. The input will include the `Client` variable, an identifier variable if needed, and a variable with included fields that match the required data to upload if needed. The return will be either a success/fail, or a variable with fields matching the site's output.

| Function | API Endpoint | Type | Input Body | Output Body |
| :------- | :----------- | :--: | :---- | :----- |
| ListBadgesForUser | /user-badges/{username}.json | GET || ListBadgesForUserResponse |
| CreateBadge | /admin/badges.json | POST | Badge | UpdatedBadgeData |
| UpdateBadgeByID | /admin/badges/{id}.json | PUT | Badge | UpdatedBadgeData |
| CreateCategory | /categories.json | POST | NewCategory | ShowCategoryResponse |
| ListCategories | /categories.json | GET || ListCategoriesResponse |
| UpdateCategoryByID | /categories/{id}.json | PUT | NewCategory | ShowCategoryResponse |
| GetCategoryContentsByID | /c/{id}.json | GET || CategoryContents |
| GetCategoryContentsBySlug | /c/{slug}.json | GET || CategoryContents |
| ShowCategory | /c/{id}/show.json | GET || ShowCategoryResponse |
| GetPersonalNotifications | /notifications.json | GET || GetNotificationsResponse |
| GetLatestPosts | /posts.json | GET || GetLatestPostsResponse |
| CreatePost | /posts.json | POST | NewPost | PostData |
| GetPostByID | /posts/{id}.json | GET || PostData |
| GetPostRepliesByID | /posts/{id}/replies.json | GET || []PostData |
| CreateTopic | /posts.json | POST | NewPost | PostData |
| GetTopicByID | /t/{id}.json | GET || TopicData |
| BookmarkTopicByID | /t/{id}/bookmark.json | PUT |||
| Search | /search.json | GET | See [Below](#the-search-function) | SearchResult |
| GetSiteInfo | /site.json | GET || SiteInfo |
| GetSiteBasicInfo | /site/basic-info.json | GET || SiteBasicInfo |
| ListTagGroups | /tag_groups.json | GET || ListTagGroupsResponse |
| GetTagGroupByID | /tag_groups/{id}.json | GET || TagGroup |
| ListTags | /tags.json | GET || ListTagsResponse |
| GetTagByName | /tag/{name}.json | GET || TagData |

#### The Search Function
Discourse sites have the special `/search.json` endpoint that allows for custom queries. This can be accessed using `Search()` with a `SearchQuery` object. Fields represent the following:

| Field | Type | Info |
| :---- | :--- | :---------- |
| Term | string | Text content to look for, can not start with `#` or `@`, can contain spaces |
| Username | string | Get posts created by this user |
| Category | string | Slug of category to limit search to |
| Tags | []string | Tags associated with each post |
| RequiresAllTags | bool | When true, posts will only show if they have all the tags provided
| Before | Time | Search for posts made before this date (not inclusive) |
| After | Time | Search for posts made on or after this date (inclusive) |
| Order | string | The order to present data in, by default you can use `OrderLatest`, `OrderLikes`, `OrderViews`, or `OrderLatestTopic` |
| AssignedUsername | string | Get items assigned to this user |
| In | []string | Where the search term is located, by default you can use `InTitle`, `InLikes`, `InPersonal`, `InMessages`, `InSeen`, `InUnseen`, `InPosted`, `InCreated`, `InWatching`, `InTracking`, `InBookmarks`, `InAssigned`, `InUnassigned`, `InFirst`, `InPinned`, or `InWiki` |
| With | []string | Only get posts containing this type of data, by default you can use `WithImages` |
| Status | []string | The post has at least one of the provided statuses, by default you can use `StatusOpen`, `StatusClosed`, `StatusPublic`, `StatusArchived`, `StatusNoReplies`, `StatusSingleUser`, `StatusSolved`,	or `StatusUnsolved` |
| Group | string | Group name or ID associated with an item |
| GroupMessages | string | Get messages associated with group name or ID |
| MinPosts | int | Minimum number of contained posts, > 0 |
| MaxPosts | int | Maximum number of contained posts, > 0 |
| MinViews | int | Minimum views on a post, > 0 |
| MaxViews | int | Maximum views on a post, > 0 |
| Custom | map[string][]string | Any additional fields specific to a Discourse site, matches the Status format of `key:val1,val2,val3` |
