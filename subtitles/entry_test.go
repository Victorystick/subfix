package subtitles

import (
	"testing"
	"time"
)

func TestEntryShift(t *testing.T) {
	e := Entry{}

	if e.start.Second() != 0 {
		t.Error("An empty Entry should have a start time of 0 secs!")
	}

	e.start = e.start.Add(time.Second)

	if e.start.Second() != 1 {
		t.Error("A second later it should be 1!")
	}
}
