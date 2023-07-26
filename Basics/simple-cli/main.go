package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var currentTask Pomodoro = Pomodoro{}

func main() {

	// currentTask = AddTask("Write a blog")
	// fmt.Println(PrintTask(currentTask))

	taskPtr := flag.String("task", "", "Task to add. (Rrquired)")
	flag.Parse()

	if *taskPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	currentTask = AddTask(*taskPtr)
	fmt.Println(PrintTask(currentTask))
	
	fmt.Printf("Start the task! Focus on %s\n", *taskPtr)

	timer1 := time.NewTimer(5 * time.Second)
	<-timer1.C

	fmt.Println("Congrats! Task time is complete. Take a break")

}