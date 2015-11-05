// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package clock_test

import (
	"time"

	_ "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/clock"
	"github.com/juju/clock/testing"
)

// Type assertions: both the wall clock and the testing clock are Clocks.
var _ clock.Clock = clock.WallClock
var _ clock.Clock = (*testing.Clock)(nil)

type clockSuite struct {
}

var _ = gc.Suite(&clockSuite{})

func (*clockSuite) now(c *gc.C) time.Time {
	utc, err := time.LoadLocation("UTC")
	c.Assert(err, gc.IsNil)
	return time.Date(2015, 9, 10, 13, 14, 15, 0, utc)
}

func (s *clockSuite) TestClockNow(c *gc.C) {
	now := s.now(c)
	tc := testing.NewClock(now)
	c.Assert(tc.Now(), gc.Equals, now)
}

func (s *clockSuite) TestClockAdvance(c *gc.C) {
	now := s.now(c)
	tc := testing.NewClock(now)
	advance := 5 * time.Minute
	tc.Advance(advance)
	c.Assert(tc.Now(), gc.Equals, now.Add(advance))
}

func (s *clockSuite) TestClockAfter(c *gc.C) {
	now := s.now(c)
	tc := testing.NewClock(now)
	advance := 5 * time.Minute

	alarm := tc.After(advance)

	select {
	case <-alarm:
		c.Error("alarm shouldn't fire yet")
	default:
		// all good
	}

	tc.Advance(advance)

	select {
	case when := <-alarm:
		c.Assert(when, gc.Equals, tc.Now())
	case <-time.After(time.Millisecond):
		c.Error("alarm should have fired")
	}
}
