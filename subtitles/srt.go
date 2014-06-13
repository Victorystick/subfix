package subtitles

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	srtTime = "15:04:05.000"
)

func ParseSrt(content string) (*Subtitles, error) {
	lines := strings.Split(content, "\n")

	subs := new(Subtitles)

	next := 1
	var entry *Entry = nil

	text := ""

	for _, line := range lines {
		if line == "" {
			if entry != nil {
				entry.frags = fragsFromText(text)
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

			entry = &Entry{id: next}
			next++
		} else if entry.start.Equal(emptyTime) {
			times := strings.Split(line, " --> ")

			t, err := time.Parse(srtTime,
				strings.Replace(times[0], ",", ".", 1))

			if err != nil {
				return nil, err
			}

			entry.start = t

			t, err = time.Parse(srtTime,
				strings.Replace(times[1], ",", ".", 1))

			if err != nil {
				return nil, err
			}

			entry.end = t
		} else if (text == "") {
			text = line
		} else {
			text += "\n" + line
		}
	}

	if entry != nil {
		entry.frags = fragsFromText(text)
		subs.Append(entry)
	}

	return subs, nil
}

func fragsFromText(text string) (frags []Fragment) {
	frags = append(frags, Fragment{
		text: text,
	})

	return
}

func (s Subtitles) Srt() string {
	strs := make([]string, len(s.entries))

	for i, e := range s.entries {
		strs[i] = e.Srt()
	}

	return strings.Join(strs, "\n")
}

func (e Entry) Srt() string {
	frags := make([]string, len(e.frags))

	for i, frag := range e.frags {
		frags[i] = frag.Srt()
	}

	str := fmt.Sprintf("%d\n%s --> %s\n%s\n",
		e.id,
		strings.Replace(
			e.start.Format(srtTime), ".", ",", 1),
		strings.Replace(
			e.end.Format(srtTime), ".", ",", 1),
		strings.Join(frags, ""))

	return str
}

func (f Fragment) Srt() string {
	text := f.text

	if f.italic {
		text = "<i>" + text + "</i>"
	}

	if f.bold {
		text = "<b>" + text + "</b>"
	}

	if f.underline {
		text = "<u>" + text + "</u>"
	}

	return text
}
