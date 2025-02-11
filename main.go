package main

import (
	"pratikshakuldeep456/task-management-system/pkg/tms"
	"time"
)

func main() {
	tm := tms.NewTaskManager()
	user1 := &tms.User{ID: 1,
		Name:  "Pratiksha",
		Email: "pratiksha@gmail.com"}
	taks1 := &tms.Task{
		Id:           1,
		Title:        "login",
		Description:  "create login page",
		DueDate:      time.Now().AddDate(2025, 2, 15),
		Priority:     tms.High,
		Status:       tms.Pending,
		AssignedUser: user1,
	}
	tm.CreateTask(taks1)

}
