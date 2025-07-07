package todo

import "encoding/json"

// Task представляет собой задачу с идентификатором, заголовком и статусом выполнения.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// AsJSON возвращает JSON-представление задачи.
func (task *Task) AsJSON() (string, error) {
	jsonData, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// NewTask создаёт новую задачу с указанными параметрами.
func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Completed: false,
	}
}
