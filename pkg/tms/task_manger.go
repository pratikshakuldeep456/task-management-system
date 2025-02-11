package tms

import (
	"strings"
	"sync"
	"time"
)

type TaskManager struct {
	Tasks map[int]*Task
	Users map[int][]*Task
	//Mu    sync.Mutex
}

var TaskManagerInstance *TaskManager
var Once sync.Once

func NewTaskManager() *TaskManager {
	Once.Do(func() {
		TaskManagerInstance = &TaskManager{
			Tasks: make(map[int]*Task),
			Users: make(map[int][]*Task),
		}
	})
	return TaskManagerInstance
}

func (tm *TaskManager) CreateTask(task *Task) {
	tm.Tasks[task.Id] = task
	//assign task to users
	tm.Users[task.AssignedUser.ID] = append(tm.Users[task.AssignedUser.ID], task)

}

func (tm *TaskManager) UpdateTask(task *Task) {
	i, exist := tm.Tasks[task.Id]
	if exist {
		i.AssignedUser = task.AssignedUser
		i.Description = task.Description
		i.DueDate = task.DueDate
		i.Priority = task.Priority
		i.Title = task.Title
		i.Status = task.Status
	}

}

func (tm *TaskManager) DeleteTask(taskID int) {
	task, exists := tm.Tasks[taskID]
	if exists {
		tasks := tm.Users[task.AssignedUser.ID]
		for i, t := range tasks {
			if t.Id == taskID {
				tm.Users[task.AssignedUser.ID] = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		//	delete(tm.Tasks)
		delete(tm.Users, taskID)
	}

}

// assign task to other user
func (tm *TaskManager) AssignTaskToUser(user *User, task *Task) {
	tm.Users[user.ID] = append(tm.Users[user.ID], task)
}

//set reminder

//systme searching and filtering tasks based on various criteria (e.g., priority,
// due date,
// assigned user).

func Contains(des, keyword string) bool {

	return strings.Contains(strings.ToLower(des), strings.ToLower(keyword))
}

func (tm *TaskManager) FilterTasks(keyword string, priority *Priority, status *Status, userID int, dueDate *time.Time) []*Task {
	var matchedTasks []*Task
	for _, task := range tm.Tasks {
		if priority != nil && task.Priority != *priority {
			continue
		}
		if status != nil && task.Status != *status {
			continue
		}
		if task.AssignedUser.ID != userID {
			continue
		}
		if dueDate != nil && !task.DueDate.Equal(*dueDate) {
			continue
		}
		if keyword != "" && !Contains(task.Description, keyword) {
			continue
		}
		matchedTasks = append(matchedTasks, task)
	}
	return matchedTasks
}

//priority *Priority, status *Status, userID int,

// users: mark task completed

func (tm *TaskManager) MarkTaskCompleted(taskid int) {
	task, exists := tm.Tasks[taskid]
	if exists {
		tm.Tasks[task.Id].Status = Completed
	}

}

func (tm *TaskManager) TaskHistory(userID int) []*Task {
	return tm.Users[userID]

}

//users: view their task history
