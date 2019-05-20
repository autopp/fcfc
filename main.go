package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Commands []Command
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

	b, err := ioutil.ReadFile(filepath.Join(home, ".fcfc.yml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var cfg Config
	yaml.Unmarshal(b, &cfg)

	return 0
}
