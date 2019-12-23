package main

import (
	"fmt"
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

	os.Exit(m.Run())
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

func TestMakeCfHome(t *testing.T) {
	c := &Command{
		Name:         "mycommand",
		API:          "api.run.pivotal.io",
		Org:          "myorg",
		Space:        "myspace",
		LoginOptions: "--sso",
	}

	cfHome := filepath.Join(homeDir, ".fcfc", "mycommand")
	expected := fmt.Sprintf(`mkdir -p "%s"`, cfHome)
	actual, err := c.MakeCfHome()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestLoginAlias(t *testing.T) {
	c := &Command{
		Name:         "mycommand",
		API:          "api.run.pivotal.io",
		Org:          "myorg",
		Space:        "myspace",
		LoginOptions: "--sso",
	}

	cfHome := filepath.Join(homeDir, ".fcfc", "mycommand")
	expected := fmt.Sprintf(`alias login-mycommand="CF_HOME=%s cf login -a api.run.pivotal.io -o myorg -s myspace --sso"`, cfHome)
	actual, err := c.LoginAlias()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCfAlias(t *testing.T) {
	c := &Command{
		Name:         "mycommand",
		API:          "api.run.pivotal.io",
		Org:          "myorg",
		Space:        "myspace",
		LoginOptions: "--sso",
	}

	cfHome := filepath.Join(homeDir, ".fcfc", "mycommand")
	expected := fmt.Sprintf(`alias mycommand="CF_HOME=%s cf"`, cfHome)
	actual, err := c.CfAlias()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
