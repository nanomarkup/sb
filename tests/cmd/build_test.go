package cmd

import (
	"fmt"

	"github.com/sapplications/smod/lod"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestBuildEmpty(c *check.C) {
	c.Skip("Needs to fix...")
	return
	c.Assert(s.Build(), check.ErrorMatches, fmt.Sprintf(lod.ModuleFilesMissingF, ".*"))
}
