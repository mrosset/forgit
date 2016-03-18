package main

import (
	"fmt"
	"github.com/str1ngs/util/file"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	FETCH_HEAD = ".git/FETCH_HEAD"
)

func main() {

	if !file.Exists(".git") {
		return
	}

	if file.Exists(".git") && !file.Exists(FETCH_HEAD) {
		git_fetch()
		return
	}

	stat, err := os.Stat(FETCH_HEAD)
	if err != nil {
		log.Fatal(err)
	}
	dur := time.Now().Sub(stat.ModTime())
	if dur > time.Minute*5 {
		fmt.Println("last fetch was ", dur)
		git_fetch()
	}
}

func git_fetch() {
	fmt.Println("fetching...")
	git := exec.Command("git", "fetch", "--all")
	git.Stderr = os.Stderr
	git.Stdout = os.Stdout
	err := git.Run()
	if err != nil {
		log.Fatal(err)
	}
}
