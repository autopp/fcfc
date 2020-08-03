/*
	Copyright (C) 2019-2020 Akira Tanimura

		Licensed under the Apache License, Version 2.0 (the "License");
		you may not use this file except in compliance with the License.
		You may obtain a copy of the License at

				http://www.apache.org/licenses/LICENSE-2.0

		Unless required by applicable law or agreed to in writing, software
		distributed under the License is distributed on an "AS IS" BASIS,
		WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
		See the License for the specific language governing permissions and
		limitations under the License.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const Version = "v0.1.0"

type Config struct {
	Commands []Command
}

func main() {
	os.Exit(run())
}

func run() int {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var configPath string
	if len(os.Args) > 0 {
		configPath = os.Args[1]
	} else {
		configPath = filepath.Join(home, ".fcfc.yml")
	}

	b, err := ioutil.ReadFile(configPath)
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
		makeCfHome, err := c.MakeCfHome()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Println(makeCfHome)

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
