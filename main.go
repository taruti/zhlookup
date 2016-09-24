// zhlookup is a utility to grep CC-CEDICT Chinese dictionary.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/chzyer/readline"
)

func main() {
	flag.Parse()
	err := mainWork()
	if err != nil {
		log.Fatal("ERROR", err)
	}
}

func mainWork() error {
	args := flag.Args()
	for _, s := range args {
		err := findPrint(s)
		if err != nil {
			return err
		}
	}

	if len(args) == 0 {
		err := console()
		if err != nil {
			return err
		}
	}

	return nil
}

func findPrint(s string) error {
	r, err := Find(s)
	if err != nil {
		return err
	}
	for _, line := range r {
		fmt.Println(line)
	}
	return nil
}

func console() error {
	rl, err := readline.New("> ")
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		err = findPrint(line)
		if err != nil { // io.EOF
			return err
		}
	}
}
