package srt

import (
	. "github.com/victorystick/subfix"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"image/color"
)

// srt.go Manages conversion to and from `.srt` subtitles.

const (
	srtTime = "15:04:05.000"
)

var emptyTime time.Time

func Register() {
	AddParser("srt", Parse)
	AddEmitter("srt", Emit)
}

func Parse(content string) (*Subtitles, error) {
	lines := strings.Split(content, "\n")

	subs := new(Subtitles)

	next := 1
	var entry *Entry = nil

	text := ""

	for _, line := range lines {
		if line == "" {
			if entry != nil {
				entry.Frags = fragsFromText(text)
				text = ""
				subs.Append(entry)
			}
			entry = nil
		} else if entry == nil {
			nr, err := strconv.Atoi(line)

			if err != nil || next != nr {
				return nil, errors.New(fmt.Sprintf(
					"Expected %d, got %s", next, line))
			}

			entry = &Entry{Id: next}
			next++
		} else if entry.Start.Equal(emptyTime) {
			times := strings.Split(line, " --> ")

			t, err := time.Parse(srtTime,
				strings.Replace(times[0], ",", ".", 1))

			if err != nil {
				return nil, err
			}

			entry.Start = t

			t, err = time.Parse(srtTime,
				strings.Replace(times[1], ",", ".", 1))

			if err != nil {
				return nil, err
			}

			entry.End = t
		} else if text == "" {
			text = line
		} else {
			text += "\n" + line
		}
	}

	if entry != nil {
		entry.Frags = fragsFromText(text)
		subs.Append(entry)
	}

	return subs, nil
}

func fragsFromText(text string) (frags []Fragment) {
	frags = append(frags, Fragment{
		Text: text,
	})

	return
}

func Emit(s *Subtitles) string {
	strs := make([]string, len(s.Entries))

	for i, e := range s.Entries {
		strs[i] = SrtEntry(e)
	}

	return strings.Join(strs, "\n")
}

func SrtEntry(e *Entry) string {
	frags := make([]string, len(e.Frags))

	for i, frag := range e.Frags {
		frags[i] = SrtFrag(frag)
	}

	str := fmt.Sprintf("%d\n%s --> %s\n%s\n",
		e.Id,
		strings.Replace(
			e.Start.Format(srtTime), ".", ",", 1),
		strings.Replace(
			e.End.Format(srtTime), ".", ",", 1),
		strings.Join(frags, ""))

	return str
}

func SrtFrag(f Fragment) string {
	text := f.Text

	if f.Italic {
		text = "<i>" + text + "</i>"
	}

	if f.Bold {
		text = "<b>" + text + "</b>"
	}

	if f.Underline {
		text = "<u>" + text + "</u>"
	}

	if f.Color != nil {
		rgba := color.RGBAModel.Convert(f.Color).(color.RGBA)

		if rgba.A != 255 {

		} else {
			text = fmt.Sprintf(
				"<font color=\"#%02x%02x%02x\">%s</font>",
				rgba.R, rgba.G, rgba.B, text)
		}

	}

	return text
}
