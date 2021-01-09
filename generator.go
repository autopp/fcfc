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

type BashGenerator struct {
}

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

	return fmt.Sprintf(`alias %s="CF_HOME=%s cf"`, c.Name, cfHome), nil
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
