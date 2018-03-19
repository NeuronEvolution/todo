package models

type TodoVisibility string

const (
	TodoVisibilityPrivate TodoVisibility = "private"
	TodoVisibilityPublic TodoVisibility = "public"
	TodoVisibilityFriend TodoVisibility = "friend"
)

type UserProfile struct {
	UserID            string
	UserName          string
	TodoVisibility TodoVisibility
}
