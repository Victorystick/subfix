package subtitles

import (
	"testing"
	"fmt"
)

func TestSrtSubEquality(t *testing.T) {
	srt, _ := ParseSrt(okSrt)
	sub, _ := ParseSub(okSub)

	// fmt.Printf("%v\n", srt)
	// fmt.Printf("%v\n", sub)

	// fmt.Println(srt.Srt())
	// fmt.Println(sub.Sub())

	// fmt.Println(srt.Sub())
	// fmt.Println(sub.Srt())

	if !srt.Equal(*sub) {
		t.Error("Equivalent subtitles should result in an equivalent representation")
	}
}

func TestSrtToSub(t *testing.T) {
	subs, err := ParseSrt(okSrt)

	if err != nil {
		t.Error(err)
	}

	if subs == nil {
		t.Error("ParseSrt shouldn't return nil after a successful parse.")
	}

	if subs.Srt() != okSrt {
		t.Error("The subtitles shoudn't be changed by parsing and printing.")
	}

	subText := subs.Sub()

	if subText != okSub {
		fmt.Printf("`%s`\n", okSub)
		fmt.Printf("`%s`\n", subText)
		t.Error("The subtitles shoudn't be changed by parsing and printing.")
	}

	subs, err = ParseSub(subText)

	if err != nil {
		t.Error(err)
	}

	if subs == nil {
		t.Error("ParseSrt shouldn't return nil after a successful parse.")
	}

	if subs.Srt() != okSrt {
		t.Error("The subtitles shoudn't be changed by parsing and printing.")
	}
}
