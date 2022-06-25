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
	sb.ModManager = &smoduleManager{Lang: lang}

	s.cmd = src.SmartBuilder{}
	s.cmd.SilentErrors = true

	s.cmd.ModManager = src.CmdManager{}
	s.cmd.ModManager.Use = "mod"
	s.cmd.ModManager.ModManager = &sb

	s.cmd.Builder = src.CmdBuilder{}
	s.cmd.Builder.Use = "build"
	s.cmd.Builder.Builder = &sb

	s.cmd.Cleaner = src.CmdCleaner{}
	s.cmd.Cleaner.Use = "clean"
	s.cmd.Cleaner.Cleaner = &sb

	s.cmd.Generator = src.CmdGenerator{}
	s.cmd.Generator.Use = "gen"
	s.cmd.Generator.Generator = &sb

	s.cmd.ModAdder = src.CmdModAdder{}
	s.cmd.ModAdder.Use = "add"
	s.cmd.ModAdder.ModManager = &sb

	s.cmd.ModDeler = src.CmdModDeler{}
	s.cmd.ModDeler.Use = "del"
	s.cmd.ModDeler.ModManager = &sb

	s.cmd.ModIniter = src.CmdModIniter{}
	s.cmd.ModIniter.Use = "init"
	s.cmd.ModIniter.ModManager = &sb
	s.cmd.Runner.SilenceErrors = true
}

func (s *CmdSuite) Mod(args ...string) error {
	setCmd("mod", args...)
	return s.cmd.Execute()
}

func (s *CmdSuite) Gen(args ...string) error {
	setCmd("gen", args...)
	return s.cmd.Execute()
}

func (s *CmdSuite) Build(args ...string) error {
	setCmd("build", args...)
	return s.cmd.Execute()
}

func (s *CmdSuite) Clean(args ...string) error {
	setCmd("clean", args...)
	return s.cmd.Execute()
}
