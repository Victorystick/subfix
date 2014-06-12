package subtitles

import (
	"image/color"
	"time"
	"strings"
)

type Subtitles struct {
	entries  []*Entry
	filename string
}

func (s *Subtitles) Append(e *Entry) {
	s.entries = append(s.entries, e)
}

func (s *Subtitles) Shift(delta time.Duration) *Subtitles {
	for _, e := range s.entries {
		e.start = e.start.Add(delta)
		e.end = e.end.Add(delta)
	}

	return s
}

type Entry struct {
	id    int
	start time.Time
	end   time.Time
	frags []Fragment
	text  []string
}

func (e *Entry) AddText(text string) {
	e.text = append(e.text, text)
}

func (e *Entry) Complete() *Entry {
	e.frags = append(e.frags, Fragment{
		text: strings.Join(e.text, "\n"),
	})

	return e
}

type Fragment struct {
	bold, italic, underline bool
	text                    string
	color                   color.Color
}

type subtitleParser func(string) (*Subtitles, error)
