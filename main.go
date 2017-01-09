// zhlookup is a utility to grep CC-CEDICT Chinese dictionary.
package main

import (
	"fmt"

	"github.com/taruti/cli"
)

type zhlookup struct{}

func main() {
	cli.Main(zhlookup{})
}

func (zhlookup) HandleCliLine(s string) error {
	return findPrint(s)
}

func (zhlookup) HandleCmdLine(ss []string) error {
	for _, s := range ss {
		err := findPrint(s)
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
