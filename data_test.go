package doomsday

import (
	"testing"
	"time"
)

type testData struct {
	t time.Time
	r time.Weekday
}

var allTestData = []testData{
	{
		t: time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC),
		r: time.Thursday,
	},
	{
		t: time.Date(2018, time.November, 2, 0, 0, 0, 0, time.UTC),
		r: time.Friday,
	},
	{
		t: time.Date(2018, time.November, 3, 0, 0, 0, 0, time.UTC),
		r: time.Saturday,
	},
	{
		t: time.Date(2018, time.November, 4, 0, 0, 0, 0, time.UTC),
		r: time.Sunday,
	},
	{
		t: time.Date(2018, time.November, 5, 0, 0, 0, 0, time.UTC),
		r: time.Monday,
	},
	{
		t: time.Date(2018, time.November, 6, 0, 0, 0, 0, time.UTC),
		r: time.Tuesday,
	},
	{
		t: time.Date(2018, time.November, 7, 0, 0, 0, 0, time.UTC),
		r: time.Wednesday,
	},
	{
		t: time.Date(2018, time.November, 8, 0, 0, 0, 0, time.UTC),
		r: time.Thursday,
	},
	{
		t: time.Date(2018, time.November, 9, 0, 0, 0, 0, time.UTC),
		r: time.Friday,
	},
	{
		t: time.Date(2018, time.November, 25, 0, 0, 0, 0, time.UTC),
		r: time.Sunday,
	},
	{
		t: time.Date(2018, time.November, 26, 0, 0, 0, 0, time.UTC),
		r: time.Monday,
	},
	{
		t: time.Date(2018, time.December, 12, 0, 0, 0, 0, time.UTC),
		r: time.Wednesday,
	},
	{
		t: time.Date(2018, time.December, 11, 0, 0, 0, 0, time.UTC),
		r: time.Tuesday,
	},
	{
		t: time.Date(2018, time.December, 9, 0, 0, 0, 0, time.UTC),
		r: time.Sunday,
	},
}

func TestGet(t *testing.T) {

	for _, iterator := range allTestData {
		result := Get(iterator.t)
		if result != iterator.r {
			t.Errorf("%s must be %s! but it %s", iterator.t, iterator.r, result)
		}
	}

}
