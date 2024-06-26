package discourse

import (
	"encoding/json"
)

type Notification struct {
	ID               int                  `json:"id"`
	UserID           int                  `json:"user_id"`
	NotificationType int                  `json:"notification_type"`
	Read             bool                 `json:"read"`
	CreatedAt        string               `json:"created_at"`
	PostNumber       int                  `json:"post_number"`
	TopicID          int                  `json:"topic_id"`
	Slug             string               `json:"slug"`
	Data             NotificationUserData `json:"data"`
}

type NotificationUserData struct {
	BadgeID    int    `json:"badge_id"`
	BadgeName  string `json:"badge_name"`
	BadgeSlug  string `json:"badge_slug"`
	BadgeTitle bool   `json:"badge_title"`
	Username   string `json:"username"`
}

type GetNotificationsResponse struct {
	Notifications          []Notification `json:"notifications"`
	TotalRowsNotifications int            `json:"total_rows_notifications"`
	SeenNotificationID     int            `json:"seen_notification_id"`
	LoadMoreNotifications  string         `json:"load_more_notifications"`
}

func GetPersonalNotifications(client *Client) (response *GetNotificationsResponse, err error) {
	data, sendErr := client.Get("notifications")

	if sendErr != nil {
		return nil, sendErr
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
