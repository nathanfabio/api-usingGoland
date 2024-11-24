package todo

import "errors"

type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

type Service struct {
	todos []Item
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (svc *Service) Add(todo string) error {
	for _, t := range svc.todos {
		if t.Task == todo {
			return errors.New("todo isn't unique")
		}
	}

	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "Created",
	})

	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todos
}
