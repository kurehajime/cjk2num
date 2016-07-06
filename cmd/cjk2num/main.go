package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/kurehajime/cjk2num"
)

func main() {
	var encode string
	var defaultEncoding string
	var text string
	var err error

	//get flags
	if runtime.GOOS == "windows" {
		defaultEncoding = "sjis"
	} else {
		defaultEncoding = "utf-8"
	}
	flag.StringVar(&encode, "e", defaultEncoding, "encoding")
	flag.Parse()

	// get str
	if len(flag.Args()) == 0 {
		text, err = readPipe()
	} else if flag.Arg(0) == "-" {
		text, err = readStdin()
	} else {
		text, err = flag.Arg(0), nil
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	text, err = transEnc(text, encode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	num, err := cjk2num.Convert(text)
	fmt.Println(fmt.Sprintf("%.0f", num))
}
