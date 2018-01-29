package models

type TodoItem struct {
	TodoID   string
	UserID   string
	Title    string
	Desc     string
	Priority int32
	Status   int32
}
