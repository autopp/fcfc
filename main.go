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

	fcfcDir := filepath.Join(home, ".fcfc")
	guardTmpl := `if [ ! -d "%s" ]; then
  if [ -e "%s" ]; then
    echo %s is already exists and is not directory >&2
  else
    mkdir -p "%s"
  fi
fi

`
	fmt.Printf(guardTmpl, fcfcDir, fcfcDir, fcfcDir, fcfcDir)
	for _, c := range cfg.Commands {
		loginAlias, err := c.LoginAlias()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Println(loginAlias)

		cfAlias, err := c.CfAlias()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Println(cfAlias)
		fmt.Println()
	}

	return 0
}
