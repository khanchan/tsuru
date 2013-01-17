// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package heal

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestRegisterAndGetHealer(c *C) {
	var h Healer
	Register("my-healer", h)
	got, err := Get("my-healer")
	c.Assert(err, IsNil)
	c.Assert(got, DeepEquals, h)
	_, err = Get("unknown-healer")
	c.Assert(err, ErrorMatches, `Unknown healer: "unknown-healer".`)
}

func (s *S) TestAll(c *C) {
	var h Healer
	Register("healer1", h)
	Register("healer2", h)
	healers := All()
	expected := map[string]Healer{
		"healer1": h,
		"healer2": h,
	}
	c.Assert(healers, DeepEquals, expected)
}
