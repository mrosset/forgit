package main

import (
	"fmt"
	"github.com/str1ngs/util/file"
)

const (
	FETCH_HEAD = ".git/FETCH_HEAD"
)

func main() {
	if !file.Exists(FETCH_HEAD) {
		return
	}

	fmt.Prinln("OK")
	//fmt.
}
