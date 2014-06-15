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
	shift   time.Duration
)

func main() {
	flag.StringVar(&outfile, "outfile", "", "Name of the output file.")
	flag.DurationVar(&shift, "ts", 0, "Amount of time to shift the subtitles.")

	flag.Parse()

	fmt.Println("args", flag.NArg())
	fmt.Println("outfile", outfile)
	fmt.Println("shift", shift)

	if flag.NArg() != 1 {
		fmt.Printf("Not enough arguments to subfix, try: %s [options] file\n", os.Args[0])
		printUsage()
		os.Exit(0)
	}

	filename := flag.Arg(0)

	if outfile == "" {
		outfile = filename
	}

	subs, err := subtitles.ReadFile(filename)

	die(err)

	if outfile != filename {
		ext, err := subtitles.Extension(outfile)

		die(err)

		if uint64(shift) != 0 {
			subs.Shift(shift)
		}

		text, err := subs.As(ext)

		die(err)

		err = ioutil.WriteFile(outfile, []byte(text), 0666)

		die(err)
	} else {
		fmt.Println(filename + " was successfully parsed.")
	}
}

func die(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage of subfix:")
	flag.PrintDefaults()
}
