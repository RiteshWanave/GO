package main

import (
	"time"
	"flag"
	"github.com/RiteshWanave/GO/mongo-todo/pkg/models"
	"runtime"
)



func main () {
	var binName string = "mongo-todo"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	task := flag.String("task", "", "Task to be added")
	complete := flag.String("complete", "", "Task to be completed")
	list := flag.Bool("list", false, "List all tasks")
	flag.Parse()

	if *task != "" {
		id := models.Add(*task)
		println("Task added")
		println("Task id is: ", id)
	}

	if *complete != "" {
		models.Complete(*complete)
		println("Task completed")
	}

	if *list {
		tasks := models.List()
		for _, task := range tasks {
			println("Task:", task.TaskName, " Completed:", task.Done, " Created At:", task.CreatedAt.Format(time.RFC822), " Completed At:", task.CompletedAt.Format(time.RFC822))
		}
	}

	time.Sleep(1 * time.Second)
}
