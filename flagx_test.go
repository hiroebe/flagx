package flagx_test

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/hiroebe/flagx"
)

func TestFlag(t *testing.T) {
	type Date struct {
		Year, Month, Day int
	}
	dateFlag := flagx.Config[Date]{
		Parse: func(s string) (Date, error) {
			tm, err := time.Parse("2006-01-02", s)
			if err != nil {
				return Date{}, err
			}
			year, month, day := tm.Date()
			return Date{Year: year, Month: int(month), Day: day}, nil
		},
		String: func(d Date) string {
			return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
		},
	}

	cases := map[string]struct {
		defaultValue Date
		args         []string
		wantErr      bool
		want         Date
	}{
		"use flag value": {
			defaultValue: Date{Year: 2023, Month: 1, Day: 1},
			args:         []string{"-date", "2023-01-02"},
			want:         Date{Year: 2023, Month: 1, Day: 2},
		},
		"use default value": {
			defaultValue: Date{Year: 2023, Month: 1, Day: 1},
			args:         []string{},
			want:         Date{Year: 2023, Month: 1, Day: 1},
		},
		"parse error": {
			defaultValue: Date{Year: 2023, Month: 1, Day: 1},
			args:         []string{"-date", "2023-1-2"},
			wantErr:      true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var got Date
			flagSet := flag.NewFlagSet("", flag.ContinueOnError)
			flagSet.Var(dateFlag.Value(&got, tc.defaultValue), "date", "")
			err := flagSet.Parse(tc.args)
			if gotErr := (err != nil); gotErr != tc.wantErr {
				t.Errorf("unexpected error: %v", err)
			}
			if err != nil {
				return
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
