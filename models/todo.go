package models

type TodoStatus string

const (
	TodoStatusOngoing   TodoStatus = "ongoing"
	TodoStatusCompleted TodoStatus = "completed"
	TodoStatusDiscard   TodoStatus = "discard"
)

type TodoItem struct {
	TodoID   string
	UserID   string
	Category string
	Title    string
	Desc     string
	Status   TodoStatus
	Priority int32
}

type TodoItemGroup struct {
	Category     string
	TodoItemList []*TodoItem
}

type CategoryInfo struct {
	Category  string
	TodoCount int64
}

type TodoItemGroupArray []*TodoItemGroup

func (array TodoItemGroupArray) Len() int {
	return len(array)
}

func (array TodoItemGroupArray) Less(i, j int) bool {
	return array[i].Category < array[j].Category
}

func (array TodoItemGroupArray) Swap(i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
