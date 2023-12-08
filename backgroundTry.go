package main

import (
	"fmt"
	"time"
)

type Job struct {
	Id     int
	Name   string
	Status string
}

var JobQue = make(chan Job, 100)

func processjob() {

	for job := range JobQue {

		fmt.Printf("job processing ID: %d\n", job.Id)

		time.Sleep(2 * time.Second)

		job.Status = "completed"
		fmt.Printf("Job ID %d completed \n", job.Id)
	}
}

func main() {

	go processjob()

	for i := 1; i <= 10; i++ {

		newjob := Job{Id: i, Name: fmt.Sprintf("job %d", i), Status: "pending"}
		JobQue <- newjob
	}

	time.Sleep(10 * time.Second)
}
