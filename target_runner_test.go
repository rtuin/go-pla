package pla_test

import (
	// "github.com/rtuin/go-pla"
	check "gopkg.in/check.v1"
)

type TargetRunnerSuite struct{}

var _ = check.Suite(&TargetRunnerSuite{})

// var findTests = []struct {
// 	name     string
// 	targets  []pla.Target
// 	expected pla.Target
// 	err      check.Checker
// }{
// 	{"foo", make([]pla.Target, 0), pla.Target{}, check.NotNil},
// 	{"Foo", []pla.Target{pla.Target{Name: "foo"}}, pla.Target{}, check.NotNil},
// 	{"foo", []pla.Target{pla.Target{Name: "foo"}}, pla.Target{Name: "foo"}, check.IsNil},
// 	{"foo", []pla.Target{pla.Target{Name: "bar"}, pla.Target{Name: "baz"}}, pla.Target{}, check.NotNil},
// 	{"foo", []pla.Target{pla.Target{Name: "oof"}, pla.Target{Name: "foo"}}, pla.Target{Name: "foo"}, check.IsNil},
// }

// func (s *TargetRunnerSuite) TestFindTargetByName(c *check.C) {
// 	for i := range findTests {
// 		target, err := pla.FindTargetByTargetName(findTests[i].name, findTests[i].targets)
// 		c.Assert(err, findTests[i].err)
// 		c.Assert(target, check.DeepEquals, findTests[i].expected)
// 	}
// }
