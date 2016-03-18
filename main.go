package main

import (
	"fmt"
	"github.com/str1ngs/util/file"
	"github.com/str1ngs/util/json"
	glog "log"
	"os"
	"os/exec"
	"time"
)

type Expandable string

func (e Expandable) String() string {
	return expand((string)(e))
}

var (
	repos  = []Expandable{"$GOPATH/src/github.com/str1ngs/forgit"}
	log    = glog.New(os.Stderr, "", glog.Lshortfile)
	expand = os.ExpandEnv
	config = Expandable("$HOME/.repos.json")
)

func init() {
	if !file.Exists(config.String()) {
		log.Println("creating example config ", config)
		err := json.Write(&repos, config.String())
		if err != nil {
			log.Fatal(err)
		}

	}
	log.Println("load config:", config)
	err := json.Read(&repos, config.String())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fetch_all()
	c := time.Tick(10 * time.Minute)
	for _ = range c {
		fetch_all()
	}

}

func fetch_all() {
	for _, r := range repos {
		log.Println("fetching", r.String())
		fetch(r.String())
	}
}

func fetch(path string) {
	fmt.Println("fetching...")
	git := exec.Command("git", "fetch", "--all")
	git.Stderr = os.Stderr
	git.Dir = path
	git.Stdout = os.Stdout
	err := git.Run()
	if err != nil {
		log.Println("error: ", err)
	}
}
