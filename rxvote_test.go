package rxvote

import "testing"
import "time"
//import "fmt"

func TestGetVoteResult(t *testing.T) {
  tests := []struct{
    input []bool
    expected bool
  }{
    {[]bool{true, true, true}, true},
    {[]bool{true, true, false}, true},
    {[]bool{true, false, false}, false},
    {[]bool{false, false, false}, false},
  }

  for _, test := range tests {
    if GetVoteResult(test.input) != test.expected {
      t.Errorf("ERROR: GetVoteResult of %q should return %q", test.input, test.expected)
    }
  }
}

func TestVote(t *testing.T) {
  ch := make(chan []bool, 1)
  timeout := time.After(4 * time.Second)
  ch <- Vote()
  select {
//    case ret := <-ch:
//      fmt.Println(ret)
    case <-ch:
    case <-timeout:
      t.Errorf("ERROR: Timeout")
  }
}
