package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), math.Pi / 2 * 3},
		{simpleTime(0, 0, 7), math.Pi / 30 * 7},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			want := tc.angle
			if got != want {
				t.Errorf("want %v radians, but got %v", want, got)
			}
		})
	}
}
func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			want := tc.point
			if !roughlyEqualPoint(want, got) {
				t.Errorf("want %v point, but got %v", want, got)
			}
		})
	}
}

func TestSecondHand(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{150, 150 - 90}},
		{simpleTime(0, 0, 30), Point{150, 150 + 90}},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := SecondHand(c.time)
			want := tc.point
			if !roughlyEqualPoint(want, got) {
				t.Errorf("want %v point, but got %v", want, got)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			want := tc.angle
			if got != want {
				t.Errorf("want %v radians, but got %v", want, got)
			}
		})
	}
}

func TestMinuteHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			want := tc.point
			if !roughlyEqualPoint(want, got) {
				t.Errorf("want %v point, but got %v", want, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 7, 0), math.Pi * 7 / (30 * 12)},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			want := tc.angle
			if got != want {
				t.Errorf("want %v radians, but got %v", want, got)
			}
		})
	}
}

func TestHourHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		tc := c
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			want := tc.point
			if !roughlyEqualPoint(want, got) {
				t.Errorf("want %v point, but got %v", want, got)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(123, time.May, 19, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
