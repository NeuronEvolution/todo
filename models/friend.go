package models

type FriendInfo struct {
	UserID            string
	UserName          string
	TodoVisibility TodoVisibility
	TodoCount         int64
}

type FriendsQuery struct {
	PageSize  int64
	PageToken string
}
