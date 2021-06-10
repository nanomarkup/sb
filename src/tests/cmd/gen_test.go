package cmd

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/smodule"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestGenerateEmpty(c *check.C) {
	c.Assert(s.Gen(), check.ErrorMatches, fmt.Sprintf(smodule.ModuleFilesMissingF, ".*$"))
}
