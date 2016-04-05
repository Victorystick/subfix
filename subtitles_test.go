package subfix

import (
	"testing"
	"time"
)

func TestEntryShift(t *testing.T) {
	e := Entry{}

	if e.Start.Second() != 0 {
		t.Error("An empty Entry should have a start time of 0 secs!")
	}

	e.Start = e.Start.Add(time.Second)

	if e.Start.Second() != 1 {
		t.Error("A second later it should be 1!")
	}
}
