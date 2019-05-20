package main

import (
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
