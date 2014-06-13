package subtitles

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
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
			start: time.Unix(int64(start/24), int64((start%24)*1000000000/24.0)),
			end:   time.Unix(int64(end/24), int64((end%24)*1000000000/24.0)),
			frags: frags,
		}

		subs.Append(entry)
		next++
	}

	return subs, nil
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
		uint64(float64(e.start.UnixNano())/(float64(time.Second)/24.0)),
		uint64(float64(e.end.UnixNano())/(float64(time.Second)/24.0)),
		strings.Join(frags, "|"))

	return str
}

func (f Fragment) Sub() string {
	return strings.Replace(f.text, "\n", "|", -1)
}
