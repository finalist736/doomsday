package doomsday

import (
	"fmt"
	"time"
)

type WeekDay int

var (
	Pn  WeekDay = 1
	Vt  WeekDay = 2
	Sr  WeekDay = 3
	Cht WeekDay = 4
	Pt  WeekDay = 5
	Sb  WeekDay = 6
	Vs  WeekDay = 7
)

func (d WeekDay) String() string {
	switch d {
	case Pn:
		return "Понедельник"
	case Vt:
		return "Вторник"
	case Sr:
		return "Среда"
	case Cht:
		return "Четверг"
	case Pt:
		return "Пятница"
	case Sb:
		return "Суббота"
	case Vs:
		return "Воскресенье"
	default:
		return "default"
	}
}

var doomsday2018 = Sr

type doomsdayDate struct {
	d int
	m int
}

var knownDoomDays []doomsdayDate

func init() {
	knownDoomDays = make([]doomsdayDate, 0)

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

	//fmt.Printf("%v\n\n", knownDoomDays)
}

func Get(d time.Time) WeekDay {
	fmt.Printf("%v\n", d)

	var min, tmp time.Duration
	var t time.Time
	min = time.Hour * 24 * 365
	var closestDoomDay doomsdayDate
	var tmin time.Time
	for _, v := range knownDoomDays {
		t = time.Date(d.Year(), time.Month(v.m), v.d, 0, 0, 0, 0, time.UTC)
		tmp = d.Sub(t)
		if tmp < 0 {
			tmp *= -1
		}
		if min > tmp {
			min = tmp
			closestDoomDay = v
			tmin = t
		}
	}
	sb := d.Sub(tmin)
	fmt.Printf("day1: %+v; duration: %v; %v\n",
		closestDoomDay, min, tmin)

	i := int(min.Hours()/24) % 7
	fmt.Printf("day2: %v\n", i)

	if sb > 0 {
		i = int(doomsday2018) + i
	} else {
		i = int(doomsday2018) - i
	}
	if i < 0 {
		i = 7 - (i * -1)
	} else if i == 0 {
		i = 7
	}
	fmt.Printf("day3: %v\n", i)
	return WeekDay(i)
}
