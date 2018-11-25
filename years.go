package doomsday

import (
	"time"
)

var doomDaysYears map[int]time.Weekday

func init() {
	// 70 = Sb
	doomDaysYears = make(map[int]time.Weekday, 136)
	start := time.Friday
	leap := 1972
	for i := 1970; i < 1970+136; i++ {

		if i == leap {
			start += 2
			leap += 4
		} else {
			start += 1
		}
		if start > 7 {
			start = start - 7
		}
		doomDaysYears[i] = start
	}
}
