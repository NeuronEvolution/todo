package models

import (
	"fmt"
	"github.com/NeuronFramework/errors"
	"unicode/utf8"
)

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

func (t *TodoItem) ValidateParams() (err error) {
	if t.Category == "" {
		return errors.InvalidParam("分类不能为空")
	}

	if utf8.RuneCountInString(t.Category) > MAX_CATEGORY_NAME_LENGTH {
		return errors.InvalidParam(fmt.Sprintf("分类名称最多%d个字符", MAX_CATEGORY_NAME_LENGTH))
	}

	if t.Title == "" {
		return errors.InvalidParam("标题不能为空")
	}

	if utf8.RuneCountInString(t.Title) > MAX_TITLE_NAME_LENGTH {
		return errors.InvalidParam(fmt.Sprintf("标题最多%d个字符", MAX_TITLE_NAME_LENGTH))
	}

	return nil
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
