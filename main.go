package main

import (
	"github.com/str1ngs/util/file"
	"github.com/str1ngs/util/json"
	"log"
	"os"
	"os/exec"
	"time"
)

type Expandable string

func (e Expandable) String() string {
	return expand((string)(e))
}

var (
	repos = []Expandable{"$GOPATH/src/github.com/str1ngs/forgit"}
	//log    = glog.New(os.Stderr, "", glog.Lshortfile)
	expand   = os.ExpandEnv
	config   = Expandable("$HOME/.repos.json")
	interval = 5 * time.Minute
	env      = []string{"HOME", "SSH_AUTH_SOCK"}
)

func chkenv() {
	pass := true
	for _, e := range env {
		v, ok := os.LookupEnv(e)
		if !ok {
			log.Printf("%s not set and is required", e)
			pass = false
		} else {
			log.Printf("%s=%s", e, v)
		}
	}
	if !pass {
		log.Fatalf("%s environment variables are required", env)
	}
}

func init() {
	chkenv()
	if !file.Exists(config.String()) {
		log.Println("creating example config ", config)
		err := json.Write(&repos, config.String())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("add your git repositories to $HOME/.repos.json. and restart forgit")
		os.Exit(0)
	}
	log.Println("load config:", config)
	err := json.Read(&repos, config.String())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fetch_all()
	log.Printf("next fetch at %s", time.Now().Add(interval))
	c := time.Tick(interval)
	for _ = range c {
		fetch_all()
		log.Printf("next fetch at %s", time.Now().Add(interval))
	}

}

func fetch_all() {
	for _, r := range repos {
		fetch(r.String())
	}
}

func fetch(path string) {
	log.Println("fetching", path)
	git := exec.Command("git", "fetch", "--all")
	git.Stderr = os.Stderr
	git.Dir = path
	//git.Stdout = os.Stdout
	err := git.Run()
	if err != nil {
		log.Println("error: ", err)
	}
}
