package main

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestCfHomeDir(t *testing.T) {
	c := &Command{
		Name: "mycommand",
	}

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	expected := filepath.Join(home, ".fcfc", "mycommand")
	actual, err := c.CfHomeDir()
	assert.NoError(t, err)
	assert.Equal(t, actual, expected)
}
