package main

import (
	"fmt"
	"os"

	hw "github.com/SDGophers/2019-05-hash-writer"
)

const HelpText = "help:\n\t%s config_file config_option root\n"

func main() {

	if len(os.Args) != 4 {
		fmt.Printf(HelpText, os.Args[0])
		os.Exit(1)
	}

	file := os.Args[1]
	opt := os.Args[2]
	root := os.Args[3]

	fd, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf, err := hw.ParseConfig(fd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = conf.Write(root, opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
