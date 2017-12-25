package conflab

import (
	"testing"
	//"time"
	"time"
)

func TestWatch(t *testing.T) {
	Server="172.30.10.161"
	Register("a")
	time.Sleep(time.Minute*2)
}

