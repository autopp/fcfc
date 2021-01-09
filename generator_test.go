package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBashGeneratorMakeCfHome(t *testing.T) {
	g := new(BashGenerator)
	c := &Command{
		Name:         "mycommand",
		API:          "api.run.pivotal.io",
		Org:          "myorg",
		Space:        "myspace",
		LoginOptions: "--sso",
	}

	cfHome := filepath.Join(homeDir, ".fcfc", "mycommand")
	expected := fmt.Sprintf(`\mkdir -p "%s"`, cfHome)
	actual, err := g.MakeCfHome(c)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestBashGeneratorLoginAlias(t *testing.T) {
	g := new(BashGenerator)
	c := &Command{
		Name:         "mycommand",
		API:          "api.run.pivotal.io",
		Org:          "myorg",
		Space:        "myspace",
		LoginOptions: "--sso",
	}

	cfHome := filepath.Join(homeDir, ".fcfc", "mycommand")
	expected := fmt.Sprintf(`alias login-mycommand="CF_HOME=%s cf login -a api.run.pivotal.io -o myorg -s myspace --sso"`, cfHome)
	actual, err := g.LoginAlias(c)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestBashGeneratorCfAlias(t *testing.T) {
	g := new(BashGenerator)
	c := &Command{
		Name:         "mycommand",
		API:          "api.run.pivotal.io",
		Org:          "myorg",
		Space:        "myspace",
		LoginOptions: "--sso",
	}

	cfHome := filepath.Join(homeDir, ".fcfc", "mycommand")
	expected := fmt.Sprintf(`alias mycommand="CF_HOME=%s cf"`, cfHome)
	actual, err := g.CfAlias(c)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
func TestBashGeneratorWithCheckingFcfcDir(t *testing.T) {
	g := new(BashGenerator)

	fcfcHome := filepath.Join(homeDir, ".fcfc")
	lines := []string{
		"echo hello",
		"echo world",
	}
	expectedFmt := `if [ ! -d "%s" -a -e "%s" ]; then
  \echo %s is already exists and is not directory >&2
else
  \mkdir -p "%s"

  echo hello
  echo world
fi
`
	expected := fmt.Sprintf(expectedFmt, fcfcHome, fcfcHome, fcfcHome, fcfcHome)
	actual, err := g.WithCheckingFcfcDir(fcfcHome, lines)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
