package sub

import (
	. "github.com/victorystick/subfix"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"image/color"
)

// sub.go Manages conversion to and from `.sub` subtitles.

var (
	timeZero = time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
)

const (
	framesASecond = int64(time.Second) / 24
)

func Register() {
	AddParser("sub", Parse)
	AddEmitter("sub", Emit)
}

func Parse(content string) (*Subtitles, error) {
	lines := strings.Split(content, "\n")

	subs := new(Subtitles)

	next := 1

	for _, line := range lines {
		if line == "" {
			continue
		}

		index := strings.Index(line, "}")

		if index == -1 {
			return nil, errors.New("Couldn't parse start frame")
		}

		start, err := strconv.Atoi(line[1:index])

		if err != nil {
			return nil, err
		}

		line = line[index+2:]

		index = strings.Index(line, "}")

		if index == -1 {
			return nil, errors.New("Couldn't parse end frame")
		}

		end, err := strconv.Atoi(line[:index])

		if err != nil {
			return nil, err
		}

		line = line[index+1:]

		textLines := strings.Split(line, "|")

		frags := make([]Fragment, len(textLines))

		for i, line := range textLines {
			frags[i] = Fragment{
				Text: line,
			}
		}

		entry := &Entry{
			Id:    next,
			Start: framesToTime(start),
			End:   framesToTime(end),
			Frags: frags,
		}

		subs.Append(entry)
		next++
	}

	return subs, nil
}

func framesToTime(number int) time.Time {
	return timeZero.Add(
		time.Second*time.Duration(number/24) +
			time.Duration((float64(number%24)/24)*1000000000))
}

func timeToFrames(t time.Time) int {
	return (((t.Hour() * 60) +
		t.Minute()*60) +
		t.Second()*24) +
		t.Nanosecond()/(1000000000/24)
}

func Emit(s *Subtitles) string {
	strs := make([]string, len(s.Entries))

	for i, e := range s.Entries {
		strs[i] = SubEntry(e)
	}

	return strings.Join(strs, "\n") + "\n"
}

func SubEntry(e *Entry) string {
	frags := make([]string, len(e.Frags))

	for i, frag := range e.Frags {
		frags[i] = SubFrag(frag)
	}

	str := fmt.Sprintf("{%d}{%d}%s",
		timeToFrames(e.Start),
		timeToFrames(e.End),
		strings.Join(frags, ""))

	return str
}

func SubFrag(f Fragment) string {
	text := f.Text

	var styles []string

	if f.Italic {
		styles = append(styles, "i")
	}

	if f.Bold {
		styles = append(styles, "b")
	}

	if f.Underline {
		styles = append(styles, "u")
	}

	if len(styles) > 0 {
		text = fmt.Sprintf("{y:%s}%s",
			strings.Join(styles, ","), text)
	}

	if f.Color != nil {
		rgba := color.RGBAModel.Convert(f.Color).(color.RGBA)

		text = fmt.Sprintf(
			"{c:$%02x%02x%02x}%s",
			rgba.B, rgba.G, rgba.R, text)

	}

	return text
}
