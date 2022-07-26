package cmd

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/smodule"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestCleanEmpty(c *check.C) {
	c.Skip("Needs to fix...")
	return
	c.Assert(s.Clean(), check.ErrorMatches, fmt.Sprintf(smodule.ModuleFilesMissingF, ".*"))
}
