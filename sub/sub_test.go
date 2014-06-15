package subtitles

import (
	"fmt"
	"testing"
	"time"
)

const okSub = `{252}{312}Elephant's Dream
{360}{432}At the left we can see...
`
const (
	s  = time.Second
	ms = time.Millisecond
)

func TestFramesToTime(t *testing.T) {
	const start = 252
	const end = 312

	if !framesToTime(start).Equal(timeZero.Add(10*s + 500*ms)) {
		t.Log(framesToTime(start))
		t.Error("Expected ")
	}

	if !framesToTime(end).Equal(timeZero.Add(13 * s)) {
		t.Log(framesToTime(end))
		t.Error("Expected ")
	}
}

func TestTimeToFrames(t *testing.T) {

}

func TestParseSub(t *testing.T) {
	subs, err := ParseSub(okSub)

	if err != nil {
		t.Error(err)
	}

	if subs == nil {
		t.Error("ParseSrt shouldn't return nil after a successful parse.")
	}

	if subs.Sub() != okSub {
		fmt.Printf("`%s`\n", okSub)
		fmt.Printf("`%s`\n", subs.Sub())
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

const italicAndGreen =
	"{c:$00ff00}{y:i}Wooh! I'm italic and green!"

func TestFragmentString(t *testing.T) {
	frag := Fragment{
		italic: true,
		text: "Wooh! I'm italic and green!",
		color: color.RGBA{0, 0xff, 0, 0xff},
	}

	if frag.Sub() != italicAndGreen {
		t.Errorf("%s should equal %s", frag.Sub(), italicAndGreen)
	}
}
