package subfix

import (
	"errors"
	"image/color"
	"time"
)

type Stringer func(*Subtitles) string

var stringers = make(map[string]Stringer)

func AddEmitter(extension string, stringer Stringer) {
	stringers[extension] = stringer
}

type Subtitles struct {
	Entries  []*Entry
	filename string
}

func (s *Subtitles) Append(e *Entry) {
	s.Entries = append(s.Entries, e)
}

// Shifts all Entries delta time.
func (s *Subtitles) Shift(delta time.Duration) *Subtitles {
	for _, e := range s.Entries {
		e.Start = e.Start.Add(delta)
		e.End = e.End.Add(delta)
	}

	return s
}

// Returns the Subtitles as a string formatted like files
// with the given extension. An error is returned if the
// extension is unknown.
func (s *Subtitles) As(ext string) (string, error) {
	fn, ok := stringers[ext]

	if !ok {
		return "", errors.New("Cannot format subtitles with extension: " + ext)
	}

	return fn(s), nil
}

// Two pairs of Subtitles are assumed to be equal,
// if all their entries are equivalent.
func (s Subtitles) Equal(s2 Subtitles) bool {
	for i, e := range s.Entries {
		if !e.Equal(*s2.Entries[i]) {
			return false
		}
	}

	return true
}

// A subtitle Entry is any amount of text displayed on screen
// within any interval. The text itself may consist of a number of
// fragments, each with different styles.
type Entry struct {
	// Id specifies the number of the entry, starting with 1
	Id int

	// Start specifies the starting time at which the subtitle
	// is to be displayed. The video is assumed to begin
	// Year 0, January, 1st 00:00:00.000000000
	Start time.Time

	// End specifies the ending time when the subtitles
	// should no longer be displayed.
	End time.Time

	// Frags is the slice of Fragments or pieces of text that
	// are to be displayed.
	Frags []Fragment
}

// Two entries are assumed to be equal if
// their ids, start times and end times,
// and fragments are equal
func (e Entry) Equal(e2 Entry) bool {
	if e.Id != e2.Id {
		return false
	}

	if !e.Start.Equal(e2.Start) {
		return false
	}

	if !e.End.Equal(e2.End) {
		return false
	}

	for i, frag := range e.Frags {
		if !frag.Equal(e2.Frags[i]) {
			return false
		}
	}

	return true
}

type Fragment struct {
	Bold, Italic, Underline bool
	Text                    string
	Color                   color.Color
}

func (f Fragment) Equal(f2 Fragment) bool {
	if f.Color == nil {
		if f2.Color != nil {
			return false
		}
	} else {
		if f2.Color == nil {
			return false
		}

		r, g, b, a := f.Color.RGBA()
		r2, g2, b2, a2 := f2.Color.RGBA()

		colorsEqual := r == r2 && g == g2 && b == b2 && a == a2

		if !colorsEqual {
			return false
		}
	}

	return f.Bold == f2.Bold &&
		f.Italic == f2.Italic &&
		f.Underline == f2.Underline &&
		f.Text == f2.Text
}
