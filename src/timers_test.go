package spec

import (
  "testing"
  "time"
  "github.com/stretchr/testify/assert"
)

func TestSetTimeout(t *testing.T) {
  start := time.Now()
  var end time.Time
  done := make(chan bool, 1)

  fn := func () {
    end = time.Now()
    done <- true
  }

  go SetTimeout(fn, 100)
  <-done

  difference := end.Sub(start) - time.Duration(100) * time.Millisecond
  assert.True(t, difference < time.Duration(5) * time.Millisecond)
}
