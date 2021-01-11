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

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const version = "HEAD"

type Config struct {
	Commands []Command
}

func main() {
	os.Exit(run())
}

func run() int {
	generators := map[string]Generator{
		"bash": new(BashGenerator),
		"zsh":  new(BashGenerator),
		"fish": new(FishGenerator),
	}

	cmd := &cobra.Command{
		Use:  "fcfc [-v] [-s shell] [config]",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if showVersion, err := cmd.Flags().GetBool("version"); err != nil {
				return err
			} else if showVersion {
				fmt.Printf("%s %s\n", cmd.Name(), version)
				return nil
			}

			shell, err := cmd.Flags().GetString("shell")
			if err != nil {
				return err
			}

			g, ok := generators[shell]
			if !ok {
				return fmt.Errorf("not supported shell %q", shell)
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			var configPath string
			if len(args) != 0 {
				configPath = args[0]
			} else {
				configPath = filepath.Join(home, ".fcfc.yml")
			}

			b, err := ioutil.ReadFile(configPath)
			if err != nil {
				return err
			}

			var cfg Config
			yaml.Unmarshal(b, &cfg)

			fcfcDir := filepath.Join(home, ".fcfc")
			lines := make([]string, 0, len(cfg.Commands)*3)
			for _, c := range cfg.Commands {
				makeCfHome, err := g.MakeCfHome(&c)
				if err != nil {
					return err
				}
				lines = append(lines, makeCfHome)

				loginAlias, err := g.LoginAlias(&c)
				if err != nil {
					return err
				}
				lines = append(lines, loginAlias)

				cfAlias, err := g.CfAlias(&c)
				if err != nil {
					return err
				}
				lines = append(lines, cfAlias)
			}

			out, err := g.WithCheckingFcfcDir(fcfcDir, lines)
			if err != nil {
				return err
			}
			fmt.Print(out)

			return nil
		},
	}
	cmd.Flags().BoolP("version", "v", false, "print version and exit")
	cmd.Flags().StringP("shell", "s", "bash", "shell type. supported: bash (default), zsh, fish")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
