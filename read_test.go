package subfix

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	AddParser("srt", func(contents string) (*Subtitles, error) {
		return &Subtitles{}, nil
	})

	_, err := ReadFile("./dream.srt")

	if err != nil {
		t.Fatal(err)
	}

	_, err = ReadFile("./missing.srt")

	if err == nil {
		t.Fatal("Should print error when opening missing file!")
	}
}

func TestExtension(t *testing.T) {
	ext, err := Extension("x.srt")

	if err != nil {
		t.Log(err)
		t.Error("Extension should find extension of 'x.srt'")
	}

	if ext != "srt" {
		t.Error("extension of 'x.srt' should be 'srt'")
	}

	ext, err = Extension("movie.eng.sub")

	if err != nil {
		t.Log(err)
		t.Error("Extension should find extension of 'movie.eng.sub'")
	}

	if ext != "sub" {
		t.Error("extension of 'movie.eng.sub' should be 'sub'")
	}
}
