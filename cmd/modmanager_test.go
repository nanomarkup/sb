package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sapplications/sb/app"
	"github.com/sapplications/smod/lod"
	"gopkg.in/check.v1"
)

func (s *CmdSuite) TestModSubcmdMissing(c *check.C) {
	c.Assert(s.Mod(), check.ErrorMatches, SubcmdMissing)
}

func (s *CmdSuite) TestModUnknownSubcmd(c *check.C) {
	c.Assert(s.Mod("test"), check.ErrorMatches, fmt.Sprintf(UnknownSubcmdF, "test"))
}

// test the init subcommand

func (s *CmdSuite) TestModInitKindMissing(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	s.Mod("init")
}

func (s *CmdSuite) TestModInit(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module
	c.Assert(s.Mod("init", app.ModKind.SB), check.IsNil)
	// read the created module
	m := lod.Manager{}
	_, err := m.ReadAll(app.ModKind.SB)
	if err != nil {
		t, _ := ioutil.ReadFile(getModuleFileName(app.DefaultModuleName))
		fmt.Print(string(t))
		c.Error(err)
	}
}

// test the add subcommand

func (s *CmdSuite) TestModAddItemMissing(c *check.C) {
	c.Assert(s.Mod("add"), check.ErrorMatches, ItemMissing)
}

func (s *CmdSuite) TestModAddModuleMissing(c *check.C) {
	c.Assert(s.Mod("add", "helloItem"), check.ErrorMatches, ModOrDepMissing)
}

func (s *CmdSuite) TestModAddEmpty(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// add a new item
	modName := "new"
	err := s.Mod("add", "helloItem", modName)
	c.Assert(err, check.IsNil)
	c.Assert(isModuleExists(modName), check.Equals, true)
	// read the created module
	m := lod.Manager{}
	_, err = m.ReadAll(app.ModKind.SB)
	if err != nil {
		t, _ := ioutil.ReadFile(getModuleFileName(modName))
		fmt.Print(string(t))
		c.Error(err)
	}
}

func (s *CmdSuite) TestModAddItem(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module
	c.Assert(s.Mod("init", app.ModKind.SB), check.IsNil)
	// add a new item use a new cmd
	cmd := CmdSuite{}
	cmd.SetUpTest(nil)
	name := "helloItem"
	err := cmd.Mod("add", name, app.DefaultModuleName)
	c.Assert(err, check.IsNil)
	// read the created module
	mod := lod.Manager{}
	r, err := mod.ReadAll(app.ModKind.SB)
	if err != nil {
		t, _ := ioutil.ReadFile(getModuleFileName(app.DefaultModuleName))
		fmt.Print(string(t))
		c.Error(err)
	} else {
		// check the added item exist
		c.Assert(r.Items()[name], check.NotNil)
	}
}

func (s *CmdSuite) TestModAddItemDependency(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module
	c.Assert(s.Mod("init", app.ModKind.SB), check.IsNil)
	// add a new dependency item (application) to the apps item
	cmd := CmdSuite{}
	cmd.SetUpTest(nil)
	name := "hello"
	resolver := "\"Hello World!\""
	err := cmd.Mod("add", lod.AppsItemName, name, resolver)
	c.Assert(err, check.IsNil)
	// TODO verify the added dependency...
}

// test the del subcommand

func (s *CmdSuite) TestModDelModuleMissing(c *check.C) {
	c.Assert(s.Mod("del", "helloItem"), check.IsNil)
}

func (s *CmdSuite) TestModDelItemMissing(c *check.C) {
	c.Assert(s.Mod("del"), check.ErrorMatches, ItemMissing)
}

func (s *CmdSuite) TestModDelItemMissing2(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module use a new cmd
	cmd := CmdSuite{}
	cmd.SetUpTest(nil)
	c.Assert(cmd.Mod("init", app.ModKind.SB), check.IsNil)
	// try to delete the missing item
	err := s.Mod("del", "helloItem")
	c.Assert(err, check.IsNil)
}

func (s *CmdSuite) TestModDelItem(c *check.C) {
	// create a temporary folder and change the current working directory
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(c.MkDir())
	// initialize a new module use a new cmd
	cmd := CmdSuite{}
	cmd.SetUpTest(nil)
	c.Assert(cmd.Mod("init", app.ModKind.SB), check.IsNil)
	// add a new item use a new cmd
	cmd = CmdSuite{}
	cmd.SetUpTest(nil)
	name := "helloItem"
	err := cmd.Mod("add", name, app.DefaultModuleName)
	c.Assert(err, check.IsNil)
	// delete the added item
	err = s.Mod("del", name)
	c.Assert(err, check.IsNil)
	// check the item does not exist
	c.Assert(isItemExists(app.ModKind.SB, name), check.Equals, false)
}

// mod del|edit|list
// NameMissing             = "\"--name\" parameter is required"
// ModuleMissing           = "\"--mod\" parameter is required"
// ResolverMissing         = "\"--resolver\" parameter is required"
// DependencyMissing       = "\"--dep\" parameter is required"
// ItemDoesNotExistF       = "\"%s\" item does not exist"
// DependencyDoesNotExistF = "\"%s\" dependency item does not exist"

//"unknown flag: --lang go"
