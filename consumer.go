package main

import (
  "time"
  "sync"
  "fmt"
)

func consumer(jobs <-chan Job, active *sync.WaitGroup) {

  /* loop until no jobs left */
  for job := range jobs {

    fmt.Printf("Sleeping for %d on job number %d\n", job.time, job.id)

    /* sleep for a job duration */
    time.Sleep(time.Second * time.Duration(job.time))
  }

  /* call P() */
  active.Done()
}

