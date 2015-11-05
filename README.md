
# clock
    import "github.com/juju/clock"

The clock package provides an interface that provides the two most common
uses of the `time` package: `Now`, and `After`. This allows tests to
provide an alternative `Clock` implementation, particularly one where the
test is in control of the time spent waiting.





## Variables
``` go
var WallClock wallClock
```
WallClock exposes wall-clock time via the Clock interface.


## func Alarm
``` go
func Alarm(c Clock, t time.Time) <-chan time.Time
```
Alarm returns a channel that will have the time sent on it at some point
after the supplied time occurs.

This is short for c.After(t.Sub(c.Now())).



## type Clock
``` go
type Clock interface {

    // Now returns the current clock time.
    Now() time.Time

    // After waits for the duration to elapse and then sends the
    // current time on the returned channel.
    After(time.Duration) <-chan time.Time
}
```
Clock provides an interface for dealing with clocks.

















- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)