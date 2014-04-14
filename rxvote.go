package rxvote

import "time"
import "math/rand"

type Randomizer interface {
  Seed(n int64)
  Intn(n int) int
}

func GetVoteResult(votes []bool) bool {
  trueCounter, falseCounter := 0, 0

  for _, vote := range votes {
    if vote == true {
      trueCounter++
    } else {
      falseCounter++
    }
  }
  return trueCounter > falseCounter
}

func Vote() (votes []bool) {
  rand.Seed(time.Now().Unix())
  /* // sequential
  for i := 0; i < 3; i++ {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
    votes = append(votes, (rand.Intn(1) == 1))
  }
  */

  // concurrency
  ch := make(chan bool)
  for i := 0; i < 3; i++ {
    go func() {
      time.Sleep(time.Duration(1 + rand.Intn(3)) * time.Second)
      ch <- (rand.Intn(2) == 1)
    }()
  }
  for i := 0; i < 3; i++ {
    votes = append(votes, <-ch)
  }
  close(ch)
  return
}
