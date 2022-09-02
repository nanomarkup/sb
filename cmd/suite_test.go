package cmd

import (
	"os"
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/sapplications/sb/plugins"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

func setCmd(cmd string, args ...string) {
	os.Args = os.Args[:1]
	os.Args = append(os.Args, cmd)
	if len(args) > 0 {
		os.Args = append(os.Args, args...)
	}
}

type CmdSuite struct {
	cmd SmartBuilder
}

var _ = check.Suite(&CmdSuite{})

func (s *CmdSuite) SetUpTest(c *check.C) {
	sb := appSmartBuilder{}
	sb.Builder = &plugins.BuilderPlugin{}
	sb.ModManager = &smoduleManager{}
	sb.PluginHandshake = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "SMART_PLUGIN",
		MagicCookieValue: "sbuilder",
	}

	s.cmd = SmartBuilder{}
	s.cmd.SilentErrors = true

	s.cmd.ModManager = CmdManager{}
	s.cmd.ModManager.Use = "mod"
	s.cmd.ModManager.ModManager = &sb

	s.cmd.Coder = CmdCoder{}
	s.cmd.Coder.Use = "code"
	s.cmd.Coder.Coder = &sb

	s.cmd.Builder = CmdBuilder{}
	s.cmd.Builder.Use = "build"
	s.cmd.Builder.Builder = &sb

	s.cmd.Cleaner = CmdCleaner{}
	s.cmd.Cleaner.Use = "clean"
	s.cmd.Cleaner.Cleaner = &sb

	s.cmd.ModAdder = CmdModAdder{}
	s.cmd.ModAdder.Use = "add"
	s.cmd.ModAdder.ModManager = &sb

	s.cmd.ModDeler = CmdModDeler{}
	s.cmd.ModDeler.Use = "del"
	s.cmd.ModDeler.ModManager = &sb

	s.cmd.ModIniter = CmdModIniter{}
	s.cmd.ModIniter.Use = "init"
	s.cmd.ModIniter.ModManager = &sb
	s.cmd.Starter.SilenceErrors = true
}

func (s *CmdSuite) Mod(args ...string) error {
	setCmd("mod", args...)
	return s.cmd.Execute()
}

func (s *CmdSuite) Code(args ...string) error {
	setCmd("code", args...)
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
