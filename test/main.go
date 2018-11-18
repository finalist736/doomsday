package main

import (
	"fmt"
	"time"

	"github.com/finalist736/doomsday"
)

// [{4 4}]
func main() {

	now := time.Date(2018, time.January, 3, 0, 0, 0, 0, time.UTC)

	d := doomsday.Get(now)
	fmt.Printf("%v\n", d)
}
