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

// Shifts all Entries delta time.
func (s *Subtitles) Shift(delta time.Duration) *Subtitles {
	for _, e := range s.entries {
		e.start = e.start.Add(delta)
		e.end = e.end.Add(delta)
	}

	return s
}

// Returns the Subtitles as a string formatted like files
// with the given extension. An error is returned if the
// extension is unknown.
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

// Two pairs of Subtitles are assumed to be equal,
// if all their entries are equivalent.
func (s Subtitles) Equal(s2 Subtitles) bool {
	for i, e := range s.entries {
		if !e.Equal(*s2.entries[i]) {
			return false
		}
	}

	return true
}

// A subtitle Entry is any amount of text displayed on screen
// within any interval. The text itself may consist of a number of
// fragments, each with different styles.
type Entry struct {
	// id specifies the number of the entry, starting with 1
	id int

	// start specifies the starting time at which the subtitle
	// is to be displayed. The video is assumed to begin
	// Year 0, January, 1st 00:00:00.000000000
	start time.Time

	// end specifies the ending time when the subtitles
	// should no longer be displayed.
	end time.Time

	// frags is the slice of Fragments or pieces of text that
	// are to be displayed.
	frags []Fragment
}

// Two entries are assumed to be equal if
// their ids, start times and end times,
// and fragments are equal
func (e Entry) Equal(e2 Entry) bool {
	if e.id != e2.id {
		return false
	}

	if !e.start.Equal(e2.start) {
		return false
	}

	if !e.end.Equal(e2.end) {
		return false
	}

	for i, frag := range e.frags {
		if !frag.Equal(e2.frags[i]) {
			return false
		}
	}

	return true
}

type Fragment struct {
	bold, italic, underline bool
	text                    string
	color                   color.Color
}

func (f Fragment) Equal(f2 Fragment) bool {
	if f.color == nil {
		if f2.color != nil {
			return false
		}
	} else {
		if f2.color == nil {
			return false
		}

		r, g, b, a := f.color.RGBA()
		r2, g2, b2, a2 := f2.color.RGBA()

		colorsEqual := r == r2 && g == g2 && b == b2 && a == a2

		if !colorsEqual {
			return false
		}
	}

	return f.bold == f2.bold &&
		f.italic == f2.italic &&
		f.underline == f2.underline &&
		f.text == f2.text
}

type subtitleParser func(string) (*Subtitles, error)
