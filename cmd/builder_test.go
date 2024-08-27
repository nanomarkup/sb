package cmd

import (
	"fmt"

	"github.com/nanomarkup/dl"
	"github.com/nanomarkup/sb"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestBuildEmpty(c *check.C) {
	c.Assert(s.Build(), check.ErrorMatches, fmt.Sprintf(dl.ModuleFilesMissingF, sb.ModKind.SB, ".*"))
}
