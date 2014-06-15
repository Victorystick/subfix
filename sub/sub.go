package subtitles

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// sub.go Manages conversion to and from `.sub` subtitles.

var (
	timeZero = time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
)

const (
	framesASecond = int64(time.Second) / 24
)

func ParseSub(content string) (*Subtitles, error) {
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
				text: line,
			}
		}

		entry := &Entry{
			id:    next,
			start: framesToTime(start),
			end:   framesToTime(end),
			frags: frags,
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

func (s Subtitles) Sub() string {
	strs := make([]string, len(s.entries))

	for i, e := range s.entries {
		strs[i] = e.Sub()
	}

	return strings.Join(strs, "\n") + "\n"
}

func (e Entry) Sub() string {
	frags := make([]string, len(e.frags))

	for i, frag := range e.frags {
		frags[i] = frag.Sub()
	}

	str := fmt.Sprintf("{%d}{%d}%s",
		timeToFrames(e.start),
		timeToFrames(e.end),
		strings.Join(frags, "|"))

	return str
}

func (f Fragment) Sub() string {
	return strings.Replace(f.text, "\n", "|", -1)
}
