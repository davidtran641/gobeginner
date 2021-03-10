package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50

	clockCenterX = 150
	clockCenterY = 150

	secondsInHalfClock = 30
	minutesInHalfClock = 30
	minutesInClock     = minutesInHalfClock * 2
	hoursInHalfClock   = 6
	hoursInClock       = hoursInHalfClock * 2
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	return makeHand(p, secondHandLength)
}

func MinuteHand(t time.Time) Point {
	p := minuteHandPoint(t)
	return makeHand(p, minuteHandLength)
}

func HourHand(t time.Time) Point {
	p := hourHandPoint(t)
	return makeHand(p, hourHandLength)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return float64(t.Second()) / secondsInHalfClock * math.Pi
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return float64(t.Minute())/minutesInHalfClock*math.Pi + secondsInRadians(t)/minutesInClock
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return float64(t.Hour()%hoursInClock)/hoursInHalfClock*math.Pi + minutesInRadians(t)/hoursInClock
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}
