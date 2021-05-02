package hcron

import (
	"fmt"
	"testing"
	"time"

	"github.com/gorhill/cronexpr"
)

// ExampleMustParse
func TestMustParse(t *testing.T) {
	tm := time.Date(2013, time.August, 31, 0, 0, 0, 0, time.UTC)
	nextTimes := cronexpr.MustParse("0 0 29 2 *").NextN(tm, 5)
	for i := range nextTimes {
		fmt.Println(nextTimes[i].Format(time.RFC1123))
		// Output:
		// Mon, 29 Feb 2016 00:00:00 UTC
		// Sat, 29 Feb 2020 00:00:00 UTC
		// Thu, 29 Feb 2024 00:00:00 UTC
		// Tue, 29 Feb 2028 00:00:00 UTC
		// Sun, 29 Feb 2032 00:00:00 UTC
	}
}
