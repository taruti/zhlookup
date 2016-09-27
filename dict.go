package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var home = os.Getenv(`HOME`)
var cedictFilename = home + `/.zhlookup.dict`

// Find performs the dictionary lookup.
func Find(s string) ([]string, error) {
	sbs := []byte(s)

	f, err := os.Open(cedictFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var res []string
	isExact := false
	for scanner.Scan() {
		bs := scanner.Bytes()
		if len(bs) == 0 || bs[0] == '#' {
			continue
		}
		var b1, b2 int
		b1 = bytes.IndexByte(bs, ' ')
		if b1 >= 0 {
			b2 = bytes.IndexAny(bs[b1+1:], " [/")
		}
		if b2 <= 0 {
			fmt.Printf("Non-standard entry: %q\n", bs)
			continue
		}
		matchIdx := bytes.Index(bs, sbs)
		if matchIdx < 0 {
			continue
		}
		if (matchIdx == 0 || isSeparator(bs[matchIdx-1])) && (matchIdx+len(sbs) == len(bs) || isSeparator(bs[matchIdx+len(sbs)])) {
			if !isExact {
				res = nil
				isExact = true
			}
			res = append(res, string(bs))
		} else if !isExact {
			res = append(res, string(bs))
		}
	}
	return res, scanner.Err()
}

func isSeparator(b byte) bool {
	return b < 0x7F && (b < 'A' || b > 'Z') && (b < 'a' || b > 'z')
}
