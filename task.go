package main

import (
	"time"
)

type Task struct {
	ID						int
	Title 				string
	Description		string
	Completed 		bool
	CreatedAt 		time.Time
}

type TaskManager struct {
	Tasks []Task
}

func (t *TaskManager) AddTask(title string, description string, completed bool) Task {
	newID := len(t.Tasks) + 1

	newTask := Task{
		ID: newID,
		Title: title,
		Description: description,
		Completed: completed,
		CreatedAt: time.Now(),
	}

	t.Tasks = append(t.Tasks, newTask)
	return newTask
}

func (t *TaskManager) GetAllTasks() []Task {
	allTasks := t.Tasks

	return allTasks
}

func (t *TaskManager) GetTaskByID(id int) (Task, bool) {
	for _, task := range t.Tasks {
		if task.ID == id {
			return task, true
		}
	}
	return Task{}, false
}

func (t *TaskManager) UpdateTask(
	id int, title string, 
	description string, completed bool,
) (Task, bool) {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Title = title
			t.Tasks[i].Description = description
			t.Tasks[i].Completed = completed

			return t.Tasks[i], true
		}
	}

	return Task{}, false
}

func (t *TaskManager) DeleteTask(id int) (Task, bool) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return task, true
		}
	}

	return Task{}, false
}
