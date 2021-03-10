package svg

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"github.com/davidtran641/gobeginner/gowithtest/math/clockface"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`

	Cirlce Circle `xml:"circle"`
	Line   []Line `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func Write(w io.Writer, t time.Time) {

	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, secondHandTag(t))
	io.WriteString(w, minuteHandTag(t))
	io.WriteString(w, hourHandTag(t))
	io.WriteString(w, svgEnd)

}

func secondHandTag(t time.Time) string {
	p := clockface.SecondHand(t)

	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X,
		p.Y)
}

func minuteHandTag(t time.Time) string {
	p := clockface.MinuteHand(t)

	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X,
		p.Y)
}

func hourHandTag(t time.Time) string {
	p := clockface.HourHand(t)

	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X,
		p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
