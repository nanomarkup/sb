package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sapplications/sbuilder/src/app"
	src "github.com/sapplications/sbuilder/src/cmd"
	"github.com/sapplications/sbuilder/src/smodule"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestDepSubcmdMissing(c *check.C) {
	c.Assert(s.Dep(), check.ErrorMatches, src.SubcmdMissing)
}

func (s *CmdSuite) TestDepUnknownSubcmd(c *check.C) {
	c.Assert(s.Dep("test"), check.ErrorMatches, fmt.Sprintf(src.UnknownSubcmdF, "test"))
}

// test the init subcommand

func (s *CmdSuite) TestDepInitGo(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module
	c.Assert(s.Dep("init", "go"), check.IsNil)
	fmt.Print("\n\n")
	// read the created module
	m := smodule.Manager{Lang: lang}
	_, err := m.ReadAll("go")
	if err != nil {
		t, _ := ioutil.ReadFile(app.DefaultModuleFileName)
		fmt.Print(string(t))
		c.Error(err)
	}
}

func (s *CmdSuite) TestDepInitLanguageMissing(c *check.C) {
	c.Assert(s.Dep("init"), check.ErrorMatches, src.LanguageMissing)
}

func (s *CmdSuite) TestDepInitLanguageIsNotSupported(c *check.C) {
	c.Assert(s.Dep("init", "delphi"), check.ErrorMatches, fmt.Sprintf(app.LanguageIsNotSupportedF, "delphi"))
}

// dep init|add|del|edit|list
// NameMissing             = "\"--name\" parameter is required"
// ModuleMissing           = "\"--mod\" parameter is required"
// LanguageMissing         = "Language parameter is required"
// ResolverMissing         = "\"--resolver\" parameter is required"
// DependencyMissing       = "\"--dep\" parameter is required"
// ItemDoesNotExistF       = "\"%s\" item does not exist"
// DependencyDoesNotExistF = "\"%s\" dependency item does not exist"

//"unknown flag: --lang go"
