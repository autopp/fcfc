package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var homeDir string

func TestMain(m *testing.M) {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestCfHomeDir(t *testing.T) {
	c := &Command{
		Name: "mycommand",
	}

	expected := filepath.Join(homeDir, ".fcfc", "mycommand")
	actual, err := c.CfHomeDir()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
