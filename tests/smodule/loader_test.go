package smodule

import (
	"fmt"

	helper "github.com/sapplications/sb/helper/hashicorp/hclog"
	"github.com/sapplications/smod/lod"
	"gopkg.in/check.v1"
)

func (s *SModuleSuite) TestLoading(c *check.C) {
	m := lod.Manager{}
	m.SetLogger(helper.NewStdOut("test", 1))
	r, e := m.ReadAll("")
	if e != nil {
		fmt.Println(e.Error())
		c.Error()
		return
	}
	c.Assert(r, check.NotNil)
}
