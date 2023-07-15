package db

import (
	"RekaKR/Rest_API_example/internal/db/model"
	"errors"
)

var Todos = []model.Todo{
	{
		ID:        "1",
		Item:      "Clean Room",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Reed Book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Record Video",
		Completed: false,
	},
}

func GetTodoById(id string) (*model.Todo, error) {
	for i, t := range Todos {
		if t.ID == id {
			return &Todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}
