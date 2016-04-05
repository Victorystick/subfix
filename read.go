package subfix

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Parser func(string) (*Subtitles, error)

var parsers = make(map[string]Parser)

func AddParser(extension string, parser Parser) {
	parsers[extension] = parser
}

func ReadFile(filename string) (*Subtitles, error) {
	ext, err := Extension(filename)

	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	parser, ok := parsers[ext]

	if !ok {
		return nil, errors.New("Unrecognized subtitle format: " + ext)
	}

	subs, err := parser(string(bytes))

	if err != nil {
		return nil, err
	}

	return subs, nil
}

func Extension(filename string) (string, error) {
	lastDot := strings.LastIndex(filename, ".")

	if lastDot == -1 {
		return "", errors.New("Unknown subtitle format. No extension found.")
	}

	ext := filename[lastDot+1:]

	return ext, nil
}
