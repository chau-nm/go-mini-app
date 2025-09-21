package main

var (
	currentID = 1
)

type Manage struct {
	todoList *[]Todo
}

func NewManage() *Manage {
	var todos []Todo
	return &Manage{
		todoList: &todos,
	}
}

func (manage *Manage) GetList() *[]Todo {
	return manage.todoList
}

func (manage *Manage) Add(task string) {
	*manage.todoList = append(*manage.todoList, Todo{
		ID:        currentID,
		Task:      task,
		Completed: false,
	})
	currentID++
}

func (manage *Manage) Complete(taskID int) {
	for i, todo := range *manage.todoList {
		if todo.ID == taskID {
			(*manage.todoList)[i].Completed = true
		}
	}
}

func (manage *Manage) Delete(taskID int) {
	for i, todo := range *manage.todoList {
		if todo.ID == taskID {
			*manage.todoList = append((*manage.todoList)[:i], (*manage.todoList)[i+1:]...)
		}
	}
}
