package conflab

import (
	"testing"
	//"time"
	"time"
)

func TestWatch(t *testing.T) {
	Register("a")
	time.Sleep(time.Minute*2)
}

