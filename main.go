package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type config struct {
	Commands []struct {
		Name         string
		API          string
		Org          string
		Space        string
		LoginOptions string `yaml:"login-options"`
	}
}

func main() {
	os.Exit(run())
}

func run() int {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	var c config
	b, err := ioutil.ReadFile(path.Join(home, ".fcfc.yml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	yaml.Unmarshal(b, &c)

	return 0
}
