/*
	Copyright (C) 2019-2021 Akira Tanimura

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
	"strings"
)

type Generator interface {
	MakeCfHome(c *Command) (string, error)
	LoginAlias(c *Command) (string, error)
	CfAlias(c *Command) (string, error)
	WithCheckingFcfcDir(fcfcDir string, lines []string) (string, error)
}

type BashGenerator struct{}

func (*BashGenerator) MakeCfHome(c *Command) (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`\mkdir -p "%s"`, cfHome), nil
}

func (*BashGenerator) LoginAlias(c *Command) (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`alias login-%s="CF_HOME=%s cf login -a %s -o %s -s %s %s"`, c.Name, cfHome, c.API, c.Org, c.Space, c.LoginOptions), nil
}

func (*BashGenerator) CfAlias(c *Command) (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`alias %s="CF_HOME=%s CF_PLUGIN_HOME=~ cf"`, c.Name, cfHome), nil
}

func (*BashGenerator) WithCheckingFcfcDir(fcfcDir string, lines []string) (string, error) {
	buf := new(strings.Builder)
	guardTmpl := `if [ ! -d "%s" -a -e "%s" ]; then
  \echo %s is already exists and is not directory >&2
else
  \mkdir -p "%s"

`
	fmt.Fprintf(buf, guardTmpl, fcfcDir, fcfcDir, fcfcDir, fcfcDir)
	for _, line := range lines {
		fmt.Fprintln(buf, "  "+line)
	}
	fmt.Fprintln(buf, "fi")

	return buf.String(), nil
}

type FishGenerator struct{}

func (*FishGenerator) MakeCfHome(c *Command) (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`command mkdir -p "%s"`, cfHome), nil
}

func (*FishGenerator) LoginAlias(c *Command) (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`alias login-%s="CF_HOME=%s cf login -a %s -o %s -s %s %s"`, c.Name, cfHome, c.API, c.Org, c.Space, c.LoginOptions), nil
}

func (*FishGenerator) CfAlias(c *Command) (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`alias %s="CF_HOME=%s CF_PLUGIN_HOME=~ cf"`, c.Name, cfHome), nil
}

func (*FishGenerator) WithCheckingFcfcDir(fcfcDir string, lines []string) (string, error) {
	buf := new(strings.Builder)
	guardTmpl := `if test ! -d "%s" -a -e "%s"
  command echo %s is already exists and is not directory >&2
else
  command mkdir -p "%s"

`
	fmt.Fprintf(buf, guardTmpl, fcfcDir, fcfcDir, fcfcDir, fcfcDir)
	for _, line := range lines {
		fmt.Fprintln(buf, "  "+line)
	}
	fmt.Fprintln(buf, "end")

	return buf.String(), nil
}
