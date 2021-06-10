package cmd

import (
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestGenerateEmpty(c *check.C) {
	if err := s.Generate(); err == nil {
		c.Fail()
	} else {
		c.Assert(err, check.ErrorMatches, "no sb files in.*$")
	}
}
