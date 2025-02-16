package main

import (
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"fmt"
	"errors"
)

func main() {
	setup()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your task (ex: CS421 HW): ")
	task, _ := reader.ReadString('\n')
	fmt.Println(task)

	fmt.Print("Enter the due date (ex: 2/20/2025): ")
	dueDate, _ := reader.ReadString('\n')
	fmt.Println(dueDate)
}

func setup() {
	// Create a new directory which contains task files
	err := os.Mkdir("tasks", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}


	if _, err := os.Stat("tasks/task.csv"); errors.Is(err, os.ErrNotExist) {
		csv_file, err := os.Create("tasks/task.csv")
		if err != nil {
			log.Fatalln("error creating a new task file", err)
		}
	} else {

	}

		w := csv.NewWriter(csv_file)


	records := [][]string {
		{"ID", "task", "status", "", "Due"},
	}

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	csv_file.Close()
}

