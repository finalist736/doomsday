package main

import (
	"fmt"
	"time"

	"github.com/finalist736/doomsday"
)

// +сменить WeekDay на time.Weekday
// +Сделать рабочим любой год
// +добавить все doomdays от 1970 + 136 лет
// +высокосный год!
func main() {

	now := time.Date(2018, time.November, 25, 0, 0, 0, 0, time.UTC)

	var d time.Weekday
	d = doomsday.Get(now)
	fmt.Printf("%v\n", d)
}
