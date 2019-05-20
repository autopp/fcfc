package main

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type Command struct {
	Name         string
	API          string
	Org          string
	Space        string
	LoginOptions string `yaml:"login-options"`
}

func (c *Command) CfHomeDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".fcfc", c.Name), nil
}

func (c *Command) LoginAlias() (string, error) {
	cfHome, err := c.CfHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`login-%s="CF_HOME=%s cf login -a %s -o %s -s %s %s"`, c.Name, cfHome, c.API, c.Org, c.Space, c.LoginOptions), nil
}
