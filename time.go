package timezh

import "time"

// Time same as time.Time, but support Chinese formatting.
type Time struct {
	time.Time
	f format
}

// T returns a Time instance from time.Time.
func T(t time.Time) Time {
	return Time{Time: t}
}

// Since returns the time elapsed since t.
// It is shorthand for time.Now().Sub(t).
func Since(t Time) time.Duration {
	return time.Now().Sub(t.Time)
}

// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
func Until(t Time) time.Duration {
	return t.Sub(time.Now())
}

// Now returns the current local time.
func Now() Time {
	return T(time.Now())
}

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(sec int64, nsec int64) Time {
	return T(time.Unix(sec, nsec))
}

// Date returns the Time corresponding to
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return T(time.Date(year, month, day, hour, min, sec, nsec, loc))
}
