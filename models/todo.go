package models

type TodoItem struct {
	TodoID   string
	UserID   string
	Category string
	Title    string
	Desc     string
	Status   string
	Priority int32
}

type TodoItemGroup struct {
	Category     string
	TodoItemList []*TodoItem
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
