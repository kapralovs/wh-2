package models

type User struct {
	ID       int64 `json:"id,omitempty"`
	RoleID   int64 `json:"role_id,omitempty"`
	JobID    int64
	Username string `json:"username,omitempty"`
}
