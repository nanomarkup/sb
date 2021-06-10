package cmd

import (
	"os"
	"testing"

	"github.com/sapplications/sbuilder/src/app"
	src "github.com/sapplications/sbuilder/src/cmd"
	"github.com/sapplications/sbuilder/src/golang"
	"github.com/sapplications/sbuilder/src/smodule"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

func lang() string {
	return "go"
}

type CmdSuite struct {
	cmd src.SmartBuilder
}

var _ = check.Suite(&CmdSuite{})

func (s *CmdSuite) SetUpSuite(c *check.C) {
	sb := app.SmartBuilder{}
	sb.Lang = lang
	sb.Manager = &smodule.Manager{Lang: lang}
	sb.GoGenerator = &golang.Generator{}
	s.cmd = src.SmartBuilder{}
	s.cmd.Generator = src.Generator{}
	s.cmd.Generator.Use = "gen"
	s.cmd.Generator.Generator = &sb
	s.cmd.Runner.SilenceErrors = true
}

func (s *CmdSuite) Generate() error {
	os.Args[1] = "gen"
	return s.cmd.Execute()
}
