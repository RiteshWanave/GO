package main

import (
	"time"
)

type Pomodoro struct {
	TaskName string
	startTime time.Time
}

func AddTask (taskname string) Pomodoro {
	return Pomodoro {
		TaskName: taskname,
		startTime: time.Now(), 
	}
}

func PrintTask (task Pomodoro) string {
	if (Pomodoro{} == task) {
		return "Current task is empty"	
	} else {
		return "Task - " + task.TaskName
	}
}

