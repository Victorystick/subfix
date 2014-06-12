package main

import (
	"github.com/victorystick/subfix/subtitles"
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"time"
)

// usage:
// subfix file delay
func main() {
	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 2 {
		printUsage()
		os.Exit(0)
	}

	filename := flag.Arg(0)

	if flag.NArg() == 1 {
		validateSubtitles(filename)

		fmt.Println(filename + " was successfully parsed.")
	} else {
		shift, err := time.ParseDuration(flag.Arg(1))

		die(err)

		subs := shiftSubtitles(filename, shift)

		ioutil.WriteFile(filename, []byte(subs.Srt()), 0)
	}
}

func validateSubtitles(filename string) {
	_, err := subtitles.ReadFile(filename)

	die(err)
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
