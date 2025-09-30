package timely

import (
	"fmt"
	"time"
)

type Now interface {
	Now() time.Time
}

type Time struct{}

func (t *Time) Now() time.Time {
	fmt.Printf("Getting current time\n")
	return time.Now()
}
