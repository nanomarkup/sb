package cmd

import (
	"fmt"

	"github.com/sapplications/dl"
	"github.com/sapplications/sb"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestBuildEmpty(c *check.C) {
	c.Assert(s.Build(), check.ErrorMatches, fmt.Sprintf(dl.ModuleFilesMissingF, sb.ModKind.SB, ".*"))
}
