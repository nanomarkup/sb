package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sapplications/sbuilder/src/app"
	"github.com/sapplications/sbuilder/src/smodule"
	"github.com/spf13/viper"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestGenEmpty(c *check.C) {
	c.Assert(s.Gen(), check.ErrorMatches, fmt.Sprintf(smodule.ModuleFilesMissingF, ".*"))
}

func (s *CmdSuite) TestGenAppMissing(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module use a new cmd
	cmd := CmdSuite{}
	cmd.SetUpTest(nil)
	c.Assert(cmd.Mod("init", lang()), check.IsNil)
	// try to generate the empty module
	c.Assert(s.Gen(), check.ErrorMatches, app.ApplicationIsMissing)
}

func (s *CmdSuite) TestGenSbApp(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	viper.Set("GOWD", wd)
	os.Chdir(c.MkDir())
	// copy sb files
	wd, _ = os.Getwd()
	copyFile("e:\\Projects\\src\\github.com\\sapplications\\sbuilder\\src\\app.sb", wd+"\\app.sb")
	copyFile("e:\\Projects\\src\\github.com\\sapplications\\sbuilder\\src\\cmd.sb", wd+"\\cmd.sb")
	copyFile("e:\\Projects\\src\\github.com\\sapplications\\sbuilder\\src\\main.sb", wd+"\\main.sb")
	// generate application's files
	c.Assert(s.Gen(), check.IsNil)
}

func (s *CmdSuite) TestGenHelloWorldApp(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	viper.Set("GOWD", wd)
	os.Chdir(c.MkDir())
	// copy sb files
	wd, _ = os.Getwd()
	copyFile("e:\\Projects\\src\\github.com\\sapplications\\sbuilder\\src\\samples\\main.sb", wd+"\\main.sb")
	// generate application's files
	c.Assert(s.Gen(), check.IsNil)
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err == nil {
		return ioutil.WriteFile(dst, input, 0644)
	} else {
		return err
	}
}
