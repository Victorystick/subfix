package subtitles

import (
	"fmt"
	"testing"
	"time"
)

const okSub = `{252}{312}Elephant's Dream
{360}{432}At the left we can see...
`

func TestParseSub(t *testing.T) {
	subs, err := ParseSub(okSub)

	if err != nil {
		t.Error(err)
	}

	if subs == nil {
		t.Error("ParseSrt shouldn't return nil after a successful parse.")
	}

	if subs.Sub() != okSub {
		fmt.Printf("`%s`", okSub)
		fmt.Printf("`%s`", subs.Sub())
		t.Error("The subtitles shoudn't be changed by parsing and printing.")
	}
}

func TestShiftSub(t *testing.T) {
	subs, err := ParseSub(okSub)

	if err != nil {
		t.Fatal(err)
	}

	if subs.entries[0].start.Second() != 10 {
		t.Error("The 0th Entry's start sec is 10")
	}

	subs.Shift(time.Second)

	if subs.entries[0].start.Second() != 11 {
		t.Error("A shift of 10s + 1s should be 11s!")
	}
}
