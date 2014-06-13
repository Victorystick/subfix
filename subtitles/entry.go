package subtitles

import (
	"errors"
	"image/color"
	"time"
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

func (s *Subtitles) As(ext string) (string, error) {
	switch ext {
	case "srt":
		return s.Srt(), nil
	case "sub":
		return s.Sub(), nil
	}

	return "",
		errors.New("Cannot format subtitles with extension: " + ext)
}

type Entry struct {
	id    int
	start time.Time
	end   time.Time
	frags []Fragment
	text  []string
}

type Fragment struct {
	bold, italic, underline bool
	text                    string
	color                   color.Color
}

type subtitleParser func(string) (*Subtitles, error)
