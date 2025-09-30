package scanner

import (
	"context"
	"fmt"
	"time"

	"github.com/jsnfwlr/timely/timely"
)

type Tool struct {
	clock timely.Now
}

func New(ctx context.Context, clock timely.Now) (tool Tool) {
	return Tool{
		clock: clock,
	}
}

type Result struct {
	ScanTime time.Time
}

func (t *Tool) Scan() (results []Result, fault error) {
	// this will be the current time, regardless of how
	// it takes to get from the top of the function to
	// this point

	res := []Result{}
	r := Result{
		ScanTime: t.clock.Now(),
	}

	fmt.Printf("Scan time: %v\n", r.ScanTime)

	res = append(res, r)

	return res, nil
}
