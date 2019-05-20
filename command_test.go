package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

var homeDir string

func TestMain(m *testing.M) {
	var err error
	homeDir, err = homedir.Dir()
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCfHomeDir(t *testing.T) {
	c := &Command{
		Name: "mycommand",
	}

	expected := filepath.Join(homeDir, ".fcfc", "mycommand")
	actual, err := c.CfHomeDir()
	assert.NoError(t, err)
	assert.Equal(t, actual, expected)
}
