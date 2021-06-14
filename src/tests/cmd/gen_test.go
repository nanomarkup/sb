package cmd

import (
	"fmt"
	"os"

	"github.com/sapplications/sbuilder/src/app"
	"github.com/sapplications/sbuilder/src/smodule"
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
