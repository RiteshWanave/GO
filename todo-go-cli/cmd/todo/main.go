package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/RiteshWanave/GO/todo-go-cli"
)

func main () {
	task := flag.String("task", "", "Task to be included in the todolist")
    list := flag.Bool("list", false, "List all tasks")
    complete := flag.Int("complete", 0, "Item to be completed")

    flag.Parse()

	l := &todo.List{}

	todoFileName := ".todo.json"

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
		
	case *task != "":
        l.Add(*task)
        if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
	
	default:
        fmt.Fprintln(os.Stderr, "Invalid option")
        os.Exit(1)
	}
}