package pla_test

import (
	"github.com/rtuin/go-pla"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type PlaTestSuite struct{}

var _ = Suite(&PlaTestSuite{})

var findTests = []struct {
	name     string
	targets  []pla.Target
	expected pla.Target
	err      Checker
}{
	{"foo", make([]pla.Target, 0), pla.Target{}, NotNil},
	{"Foo", []pla.Target{pla.Target{Name: "foo"}}, pla.Target{}, NotNil},
	{"foo", []pla.Target{pla.Target{Name: "foo"}}, pla.Target{Name: "foo"}, IsNil},
	{"foo", []pla.Target{pla.Target{Name: "bar"}, pla.Target{Name: "baz"}}, pla.Target{}, NotNil},
	{"foo", []pla.Target{pla.Target{Name: "oof"}, pla.Target{Name: "foo"}}, pla.Target{Name: "foo"}, IsNil},
}

func (s *PlaTestSuite) TestFindTargetByName(c *C) {
	for i := range findTests {
		target, err := pla.FindTargetByTargetName(findTests[i].name, findTests[i].targets)
		c.Assert(err, findTests[i].err)
		c.Assert(target, DeepEquals, findTests[i].expected)
	}
}
