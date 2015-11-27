package pla_test

import (
	"github.com/rtuin/go-pla"
	check "gopkg.in/check.v1"
)

type TargetSuite struct{}

var _ = check.Suite(&TargetSuite{})

var runTests = []struct {
	target pla.Target
	params []string
	rVal   bool
}{
	{pla.Target{Commands: make([]pla.Runnable, 0)}, make([]string, 0), false},
	{pla.Target{Commands: ([]pla.Runnable{pla.Command{"echo 1"}})}, make([]string, 0), false},
	{pla.Target{Commands: ([]pla.Runnable{pla.Command{"exit 1"}})}, make([]string, 0), true},
	{
		pla.Target{Commands: []pla.Runnable{pla.Target{Commands: []pla.Runnable{pla.Command{"echo foo"}}}}},
		make([]string, 0),
		false,
	},
	{
		pla.Target{Commands: []pla.Runnable{pla.Target{Commands: []pla.Runnable{pla.Command{"exit 1"}}}}},
		make([]string, 0),
		true,
	},
}

func (s *TargetSuite) TestRunTarget(c *check.C) {
	for i := range runTests {
		c.Assert(runTests[i].target.Run(runTests[i].params, false), check.Equals, runTests[i].rVal)
	}
}
