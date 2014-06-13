package main

import (
	"flag"
	"fmt"
	"github.com/victorystick/subfix/subtitles"
	"io/ioutil"
	"os"
	"time"
)

var (
	outfile string
)

// usage:
// subfix file delay
func main() {
	flag.StringVar(&outfile, "outfile", "", "Name ouf the output file.")

	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 2 {
		printUsage()
		os.Exit(0)
	}

	filename := flag.Arg(0)

	if outfile == "" {
		outfile = filename
	}

	if flag.NArg() == 1 {
		subs, err := subtitles.ReadFile(filename)

		die(err)

		if outfile != filename {
			ext, err := subtitles.Extension(outfile)

			die(err)

			text, err := subs.As(ext)

			die(err)

			err = ioutil.WriteFile(outfile, []byte(text), 0666)

			die(err)
		} else {
			fmt.Println(filename + " was successfully parsed.")
		}
	} else {
		ext, err := subtitles.Extension(outfile)

		die(err)

		shift, err := time.ParseDuration(flag.Arg(1))

		die(err)

		subs := shiftSubtitles(filename, shift)

		text, err := subs.As(ext)

		die(err)

		err = ioutil.WriteFile(outfile, []byte(text), 0666)

		die(err)
	}
}

func shiftSubtitles(filename string, shift time.Duration) *subtitles.Subtitles {
	subs, err := subtitles.ReadFile(filename)

	die(err)

	subs.Shift(shift)

	return subs
}

func die(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("subfix filename delay")
	fmt.Println("\tsupported filetypes:   srt")
	fmt.Println("\texample delays:        4.3s 1200ms")
}
