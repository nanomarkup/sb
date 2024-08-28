package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/nanomarkup/dl"
	"github.com/nanomarkup/sb"
	"github.com/nanomarkup/sb/plugins"
	"github.com/spf13/cobra"
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
	sc := sb.SmartCreator{}
	manager := &smoduleManager{}
	manager.Kind = sb.ModKind.SA
	sc.ModManager = manager

	sb := appSmartBuilder{}
	sb.Builder = &plugins.BuilderPlugin{}
	manager = &smoduleManager{}
	manager.Kind = "sb"
	sb.ModManager = manager
	sb.PluginHandshake = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "SMART_PLUGIN",
		MagicCookieValue: "sbuilder",
	}

	f := dl.Formatter{}

	s.cmd = SmartBuilder{}
	s.cmd.SilenceErrors = true

	modManager := cobra.Command{}
	modManager.Use = "mod"
	modManager.RunE = CmdManageMod(&sb.SmartBuilder, &f)
	s.cmd.AddCommand(&modManager)

	modAdder := cobra.Command{}
	modAdder.Use = "add"
	modAdder.RunE = CmdAddToMod(&sb.SmartBuilder)
	modManager.AddCommand(&modAdder)

	modDeler := cobra.Command{}
	modDeler.Use = "del"
	modDeler.RunE = CmdDelFromMod(&sb.SmartBuilder)
	modManager.AddCommand(&modDeler)

	modIniter := cobra.Command{}
	modIniter.Use = "init"
	modIniter.RunE = CmdInitMod(&sb.SmartBuilder)
	modManager.AddCommand(&modIniter)

	creator := cobra.Command{}
	creator.Use = "new"
	creator.RunE = CmdCreate(&sc)
	s.cmd.AddCommand(&creator)

	coder := cobra.Command{}
	coder.Use = "code"
	coder.RunE = CmdCode(&sb.SmartBuilder)
	s.cmd.AddCommand(&coder)

	builder := cobra.Command{}
	builder.Use = "build"
	builder.RunE = CmdBuild(&sb.SmartBuilder)
	s.cmd.AddCommand(&builder)

	cleaner := cobra.Command{}
	cleaner.Use = "clean"
	cleaner.RunE = CmdClean(&sb.SmartBuilder)
	s.cmd.AddCommand(&cleaner)
}

func (s *CmdSuite) Mod(args ...string) error {
	setCmd("mod", args...)
	return s.cmd.Execute()
}

func (s *CmdSuite) New(args ...string) error {
	setCmd("new", args...)
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

func getModuleFileName(name string) string {
	moduleExt := ".sb"
	if strings.HasSuffix(name, moduleExt) {
		return name
	} else {
		return name + moduleExt
	}
}

func isItemExists(kind, item string) (found bool) {
	m := dl.Manager{}
	m.Kind = kind
	all, err := m.ReadAll()
	if err != nil {
		return false
	}
	_, found = all.Items()[item]
	return
}

func isModuleExists(name string) bool {
	_, err := os.Stat(getModuleFileName(name))
	return err == nil
}
