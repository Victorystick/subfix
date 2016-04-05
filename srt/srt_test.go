package srt

import (
	. "github.com/victorystick/subfix"
	"testing"
	"time"
	"image/color"
)

const okSrt = `1
00:00:10,500 --> 00:00:13,000
Elephant's Dream

2
00:00:15,000 --> 00:00:18,000
At the left we can see...
`

func TestParseSrt(t *testing.T) {
	subs, err := Parse(okSrt)

	if err != nil {
		t.Error(err)
	}

	if subs == nil {
		t.Error("ParseSrt shouldn't return nil after a successful parse.")
	}

	if Emit(subs) != okSrt {
		t.Error("The subtitles shoudn't be changed by parsing and printing.")
	}
}

func TestShiftSrt(t *testing.T) {
	subs, err := Parse(okSrt)

	if err != nil {
		t.Error(err)
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
	"<font color=\"#00ff00\"><i>Wooh! I'm italic and green!</i></font>"

func TestFragmentString(t *testing.T) {
	frag := Fragment{
		Italic: true,
		Text: "Wooh! I'm italic and green!",
		Color: color.RGBA{0, 0xff, 0, 0xff},
	}

	if SrtFrag(frag) != italicAndGreen {
		t.Errorf("%s should equal %s", SrtFrag(frag), italicAndGreen)
	}
}
