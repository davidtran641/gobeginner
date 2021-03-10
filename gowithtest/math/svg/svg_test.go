package svg

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

func TestSVGWriteSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpletime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpletime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			Write(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(tc.line, svg.Line) {
				t.Errorf("Expected to find second hand with %v, got :\n%v", tc.line, svg)
			}
		})
	}
}

func TestSVGWriteMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpletime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			Write(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(tc.line, svg.Line) {
				t.Errorf("Expected to find minute hand with %v, got :\n%v", tc.line, svg)
			}
		})
	}
}

func TestSVGWriteHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpletime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			Write(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(tc.line, svg.Line) {
				t.Errorf("Expected to find hour hand with %v, got :\n%v", tc.line, svg)
			}
		})
	}
}

func simpletime(hours, minutes, seconds int) time.Time {
	return time.Date(123, time.May, 19, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func containsLine(line Line, lines []Line) bool {
	for _, v := range lines {
		if v == line {
			return true
		}
	}
	return false
}
