package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sapplications/sbuilder/src/app"
	"github.com/sapplications/smod/lod"
	"github.com/spf13/viper"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestGenEmpty(c *check.C) {
	c.Skip("Needs to fix...")
	return
	c.Assert(s.Gen(), check.ErrorMatches, fmt.Sprintf(lod.ModuleFilesMissingF, ".*"))
}

func (s *CmdSuite) TestGenAppMissing(c *check.C) {
	c.Skip("Needs to fix...")
	return
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module use a new cmd
	cmd := CmdSuite{}
	cmd.SetUpTest(nil)
	c.Assert(cmd.Mod("init", modType.sb), check.IsNil)
	// try to generate the empty module
	c.Assert(s.Gen(), check.ErrorMatches, app.ApplicationIsMissing)
}

func (s *CmdSuite) TestGenSbApp(c *check.C) {
	c.Skip("Needs to fix...")
	return
	// create a temporary folder and change the current working directory
	currFolder, _ := os.Getwd()
	defer os.Chdir(currFolder)
	viper.Set("GOWD", currFolder)
	os.Chdir(c.MkDir())
	// copy sb files
	wd, _ := os.Getwd()
	copyFile(currFolder+"\\..\\..\\apps.sb", wd+"\\apps.sb")
	copyFile(currFolder+"\\..\\..\\sb.sb", wd+"\\sb.sb")
	copyFile(currFolder+"\\..\\..\\sgo.sb", wd+"\\sgo.sb")
	// generate application's files
	c.Assert(s.Gen("sb"), check.IsNil)
}

func (s *CmdSuite) TestGenHelloWorldApp(c *check.C) {
	c.Skip("Needs to fix...")
	return
	// create a temporary folder and change the current working directory
	currFolder, _ := os.Getwd()
	defer os.Chdir(currFolder)
	viper.Set("GOWD", currFolder)
	os.Chdir(c.MkDir())
	// copy sb files
	wd, _ := os.Getwd()
	copyFile(currFolder+"\\..\\..\\samples\\app.sb", wd+"\\app.sb")
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
