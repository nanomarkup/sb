package cmd

import (
	"os"
	"testing"

	src "github.com/sapplications/sbuilder/src/cmd"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

func lang() string {
	return "go"
}

func setCmd(cmd string, args ...string) {
	os.Args = os.Args[:1]
	os.Args = append(os.Args, cmd)
	if len(args) > 0 {
		os.Args = append(os.Args, args...)
	}
}

type CmdSuite struct {
	cmd src.SmartBuilder
}

var _ = check.Suite(&CmdSuite{})

func (s *CmdSuite) SetUpTest(c *check.C) {
	sb := appSmartBuilder{}
	sb.Lang = lang
	sb.Manager = &smoduleManager{Lang: lang}
	sb.GoGenerator = &golangGenerator{}

	s.cmd = src.SmartBuilder{}
	s.cmd.SilentErrors = true

	s.cmd.Manager = src.Manager{}
	s.cmd.Manager.Use = "mod"
	s.cmd.Manager.Manager = &sb

	s.cmd.Builder = src.Builder{}
	s.cmd.Builder.Use = "build"
	s.cmd.Builder.Builder = &sb

	s.cmd.Cleaner = src.Cleaner{}
	s.cmd.Cleaner.Use = "clean"
	s.cmd.Cleaner.Cleaner = &sb

	s.cmd.Generator = src.Generator{}
	s.cmd.Generator.Use = "gen"
	s.cmd.Generator.Generator = &sb

	s.cmd.ModAdder = src.ModAdder{}
	s.cmd.ModAdder.Use = "add"
	s.cmd.ModAdder.Manager = &sb

	s.cmd.ModDeler = src.ModDeler{}
	s.cmd.ModDeler.Use = "del"
	s.cmd.ModDeler.Manager = &sb

	s.cmd.ModIniter = src.ModIniter{}
	s.cmd.ModIniter.Use = "init"
	s.cmd.ModIniter.Manager = &sb
	s.cmd.Runner.SilenceErrors = true
}

func (s *CmdSuite) Mod(args ...string) error {
	setCmd("mod", args...)
	return s.cmd.Execute()
}

func (s *CmdSuite) Gen() error {
	setCmd("gen")
	return s.cmd.Execute()
}

func (s *CmdSuite) Build() error {
	setCmd("build")
	return s.cmd.Execute()
}

func (s *CmdSuite) Clean() error {
	setCmd("clean")
	return s.cmd.Execute()
}
