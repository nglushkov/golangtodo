package todo

import "encoding/json"

// TaskList представляет собой список задач.
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// NewTaskList создаёт новый пустой список задач.
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: []Task{},
	}
}

// Add добавляет задачу в список.
func (list *TaskList) Add(task Task) {
	list.Tasks = append(list.Tasks, task)
}

// Get возвращает задачу по ID или nil, если задача не найдена.
func (list *TaskList) Get(id int) *Task {
	for i := range list.Tasks {
		if list.Tasks[i].ID == id {
			return &list.Tasks[i]
		}
	}
	return nil
}

// Remove удаляет задачу по ID.
func (list *TaskList) Remove(id int) {
	for i, task := range list.Tasks {
		if task.ID == id {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			return
		}
	}
}

// UpdateTask обновляет заголовок и статус задачи по ID.
func (list *TaskList) UpdateTask(id int, title string, completed bool) {
	for i, task := range list.Tasks {
		if task.ID == id {
			list.Tasks[i].Title = title
			list.Tasks[i].Completed = completed
			return
		}
	}
}

// Clear удаляет все задачи из списка.
func (list *TaskList) Clear() {
	list.Tasks = []Task{}
}

// GetAll возвращает все задачи из списка.
func (list *TaskList) GetAll() []Task {
	return list.Tasks
}

// GetAsJSON возвращает JSON-представление списка задач.
func (list *TaskList) GetAsJSON() (string, error) {
	jsonData, err := json.Marshal(list)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
