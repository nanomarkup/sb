package golang

import (
	"github.com/sapplications/sbuilder/src/golang"
	"gopkg.in/check.v1"
)

func (s *GoSuite) TestGenerate(c *check.C) {
	g := golang.Generator{}
	g.Init(
		map[string]map[string]string{
			"main": {"test": "github.com/sapplications/sbuilder/src/tests/golang.Item1"},
			"github.com/sapplications/sbuilder/src/tests/golang.Item1": {
				"Field1":    "github.com/sapplications/sbuilder/src/tests/golang.NewField1()",
				"Field2":    "github.com/sapplications/sbuilder/src/tests/golang.NewField2(\"Vitalii\")",
				"Field3":    "github.com/sapplications/sbuilder/src/tests/golang.NewField3(github.com/sapplications/sbuilder/src/tests/golang.Field1)",
				"Runner":    "*github.com/sapplications/sbuilder/src/tests/golang.RunnerImpl",
				"Hello":     "github.com/sapplications/sbuilder/src/tests/golang.Hello()",
				"EmptyFunc": "github.com/sapplications/sbuilder/src/tests/golang.EmptyFunc()"},
		},
	)
	err := g.Generate("test")
	c.Assert(err, check.IsNil)
}
