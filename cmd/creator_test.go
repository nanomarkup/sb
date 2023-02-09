// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package cmd

import (
	"os"

	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestNewSbApp(c *check.C) {
	// create a temporary folder and change the current working directory
	currFolder, _ := os.Getwd()
	defer os.Chdir(currFolder)
	os.Chdir(c.MkDir())
	// generate a new application unit
	c.Assert(s.New("sb"), check.IsNil)
}
