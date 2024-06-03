package discourse

type Group struct {
}

type GroupPermission struct {
	PermissionType int    `json:"permission_type"`
	GroupName      string `json:"group_name"`
}
