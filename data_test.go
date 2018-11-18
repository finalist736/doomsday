package doomsday

import (
	"testing"
	"time"
)

type testData struct {
	t time.Time
	r WeekDay
}

var allTestData = []testData{
	{
		t: time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC),
		r: Cht,
	},
	{
		t: time.Date(2018, time.November, 2, 0, 0, 0, 0, time.UTC),
		r: Pt,
	},
	{
		t: time.Date(2018, time.November, 3, 0, 0, 0, 0, time.UTC),
		r: Sb,
	},
	{
		t: time.Date(2018, time.November, 4, 0, 0, 0, 0, time.UTC),
		r: Vs,
	},
	{
		t: time.Date(2018, time.November, 5, 0, 0, 0, 0, time.UTC),
		r: Pn,
	},
	{
		t: time.Date(2018, time.November, 6, 0, 0, 0, 0, time.UTC),
		r: Vt,
	},
	{
		t: time.Date(2018, time.November, 7, 0, 0, 0, 0, time.UTC),
		r: Sr,
	},
	{
		t: time.Date(2018, time.November, 8, 0, 0, 0, 0, time.UTC),
		r: Cht,
	},
	{
		t: time.Date(2018, time.November, 9, 0, 0, 0, 0, time.UTC),
		r: Pt,
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
