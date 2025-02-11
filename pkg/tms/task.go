package tms

import "time"

type Priority int
type Status string

const (
	High Priority = 0
	Medium
	Low
)
const (
	Pending   Status = "pending"
	Progress  Status = "in_progress"
	Completed Status = "completed"
)

type Task struct {
	Id           int
	Title        string
	Description  string
	DueDate      time.Time
	Priority     Priority
	Status       Status
	AssignedUser *User
}

func NewTask(id int, title, description string, dueDate time.Time, priority Priority, assignedUser *User) *Task {
	return &Task{Id: id,
		Title:        title,
		Description:  description,
		DueDate:      dueDate,
		Priority:     priority,
		Status:       Pending,
		AssignedUser: assignedUser}
}
