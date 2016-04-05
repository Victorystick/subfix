package sub

import (
	. "github.com/victorystick/subfix"
	"fmt"
	"testing"
	"time"
	"image/color"
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
	subs, err := Parse(okSub)

	if err != nil {
		t.Error(err)
	}

	if subs == nil {
		t.Error("ParseSrt shouldn't return nil after a successful parse.")
	}

	if Emit(subs) != okSub {
		fmt.Printf("`%s`\n", okSub)
		fmt.Printf("`%s`\n", Emit(subs))
		t.Error("The subtitles shoudn't be changed by parsing and printing.")
	}
}

func TestShiftSub(t *testing.T) {
	subs, err := Parse(okSub)

	if err != nil {
		t.Fatal(err)
	}

	if subs.Entries[0].Start.Second() != 10 {
		t.Error("The 0th Entry's start sec is 10")
	}

	subs.Shift(time.Second)

	if subs.Entries[0].Start.Second() != 11 {
		t.Error("A shift of 10s + 1s should be 11s!")
	}
}

const italicAndGreen =
	"{c:$00ff00}{y:i}Wooh! I'm italic and green!"

func TestFragmentString(t *testing.T) {
	frag := Fragment{
		Italic: true,
		Text: "Wooh! I'm italic and green!",
		Color: color.RGBA{0, 0xff, 0, 0xff},
	}

	if SubFrag(frag) != italicAndGreen {
		t.Errorf("%s should equal %s", SubFrag(frag), italicAndGreen)
	}
}
