package main

import (
  "sync"
  "math/rand"
)

func producer(number int, jobs chan<- Job, active *sync.WaitGroup) {

  /* create <number> amout of jobs */
  for i := 0; i < number; i++ {

    /* generate random time from 2 - 4 */
    time := rand.Intn(2) + 3

    jobs <- Job{i+1, time}
  }

  /* close channel */
  close(jobs)

  /* call P() */
  defer active.Done()
}
