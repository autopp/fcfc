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
	"path/filepath"

	"os"
)

type Command struct {
	Name         string
	API          string
	Org          string
	Space        string
	LoginOptions string `yaml:"login-options"`
}

func (c *Command) CfHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".fcfc", c.Name), nil
}

func (c *Command) MakeCfHome() (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`\mkdir -p "%s"`, cfHome), nil
}

func (c *Command) LoginAlias() (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`alias login-%s="CF_HOME=%s cf login -a %s -o %s -s %s %s"`, c.Name, cfHome, c.API, c.Org, c.Space, c.LoginOptions), nil
}

func (c *Command) CfAlias() (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`alias %s="CF_HOME=%s cf"`, c.Name, cfHome), nil
}
