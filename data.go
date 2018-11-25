package doomsday

import (
	"strconv"
	"time"
)

type doomsdayDate struct {
	d int
	m int
}

var knownDoomDays []doomsdayDate

func init() {

	//ponedelnik := time.Monday

	knownDoomDays = make([]doomsdayDate, 0, 13)

	dd := doomsdayDate{d: 4, m: 4}
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 6
	dd.m = 6
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 8
	dd.m = 8
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 10
	dd.m = 10
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 12
	dd.m = 12
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 9
	dd.m = 5
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 11
	dd.m = 7
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 5
	dd.m = 9
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 7
	dd.m = 11
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 7
	dd.m = 3
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 14
	dd.m = 3
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 21
	dd.m = 3
	knownDoomDays = append(knownDoomDays, dd)

	dd.d = 28
	dd.m = 3
	knownDoomDays = append(knownDoomDays, dd)

	//fmt.Printf("init data: %v\n\n", knownDoomDays)
}

func Get(d time.Time) time.Weekday {
	//fmt.Printf("%v\n", d)

	var min, tmp time.Duration
	var t time.Time
	min = time.Hour * 24 * 365
	//var closestDoomDay doomsdayDate
	var tmin time.Time
	for _, v := range knownDoomDays {
		t = time.Date(d.Year(), time.Month(v.m), v.d, 0, 0, 0, 0, time.UTC)
		tmp = d.Sub(t)
		if tmp < 0 {
			tmp *= -1
		}
		if min > tmp {
			min = tmp
			//closestDoomDay = v
			tmin = t
		}
	}
	sb := d.Sub(tmin)
	//fmt.Printf("day1: %+v; duration: %v; %v\n",
	//	closestDoomDay, min, tmin)

	i := int(min.Hours()/24) % 7
	//fmt.Printf("day2: %v\n", i)

	doomsdayInYear, exists := doomDaysYears[d.Year()]
	if !exists {
		panic("no data for this year = " +
			strconv.FormatInt(int64(d.Year()), 10))
	}

	if sb > 0 {
		i = int(doomsdayInYear) + i
	} else {
		i = int(doomsdayInYear) - i
	}
	if i < 0 {
		i = 7 - (i * -1)
	} else if i == 7 {
		i = 0
	}
	//fmt.Printf("day3: %v\n", i)
	return time.Weekday(i)
}
