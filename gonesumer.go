package main

import (
  "math/rand"
  "time"
  "sync"
)

/* structure for job */
type Job struct {
  id int
  time int
}

func main() {

  /* WaitGroup works like a counting semaphore */
  var active sync.WaitGroup

  /* seed random number generator */
  rand.Seed(time.Now().Unix())

  /* channel that buffers jobs; acts like a queue */
  jobs := make(chan Job, 5)

  /* call V() */
  active.Add(1)
  /* run producer as a co-routine */
  go producer(5, jobs, &active)

  /* call V() */
  active.Add(1)
  /* run consumer as a co-routine */
  go consumer(jobs, &active)

  /* wait until counting semaphore is at 0 */
  active.Wait()
}
