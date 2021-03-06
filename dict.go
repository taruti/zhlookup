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
	sepQuery := len(s)>0 && isSeparator(s[0]) && isSeparator(s[len(s)-1])
	sbs := []byte(s)
	isLatinAlpha := isLatinAlphaString(sbs)

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
		if isLatinAlpha {
			toLowerLatin(bs)
		}
		matchIdx := bytes.Index(bs, sbs)
		if matchIdx < 0 {
			continue
		}
		if (matchIdx == 0 || isSeparator(bs[matchIdx-1])) && (matchIdx+len(sbs) == len(bs) || isSeparator(bs[matchIdx+len(sbs)])) {
			if !isExact && !sepQuery {
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

func isLatinAlphaString(bs []byte) bool {
	for _,b := range bs {
		if b > 0x7F || (b >= 'A' && b <= 'Z') {
			return false
		}
	}
	return true
}

func toLowerLatin(bs []byte) {
	for i,b := range bs {
		if b >= 'A' && b <= 'Z' {
			bs[i]  += 'a'-'A'
		}
	}
}

func isSeparator(b byte) bool {
	return b < 0x7F && (b < 'A' || b > 'Z') && (b < 'a' || b > 'z')
}
