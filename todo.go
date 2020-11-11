package todo

import "errors"

// TodoList model
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ListsItem struct {
	Id     int
	ItemId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Description == nil && i.Title == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
