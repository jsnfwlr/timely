package scanner_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jsnfwlr/timely/scanner"
)

type mockNow struct {
	time string
}

func (m *mockNow) Now() time.Time {
	fmt.Printf("Getting mocked time: %s\n", m.time)
	t, _ := time.Parse(time.RFC3339, m.time)
	return t
}

func TestScan(t *testing.T) {
	testCases := []struct {
		name  string
		clock *mockNow
	}{
		{
			name: "Christmas 2024",
			clock: &mockNow{
				time: "2024-12-25T12:00:00+08:00",
			},
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tool := scanner.New(ctx, tc.clock)

			res, err := tool.Scan()
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			for _, r := range res {
				if r.ScanTime != tc.clock.Now() {
					t.Errorf("expected scan time %v, got %v", tc.clock.Now(), r.ScanTime)
				}
			}
		})
	}
}
