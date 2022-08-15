package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of Go Lang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// The argument in the Format is from the documentaion.

	createdDate := time.Date(2023, time.July, 05, 16, 00, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}

/*
Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time
Date returns the Time corresponding to
	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
in the appropriate zone for that time in the given location.

The month, day, hour, min, sec, and nsec values may be outside
their usual ranges and will be normalized during the conversion.
For example, October 32 converts to November 1.

A daylight savings time transition skips or repeats times.
For example, in the United States, March 13, 2011 2:15am never occurred,
while November 6, 2011 1:15am occurred twice. In such cases, the
choice of time zone, and therefore the time, is not well-defined.
Date returns a time that is correct in one of the two zones involved
in the transition, but it does not guarantee which.

Date panics if loc is nil.


*/
