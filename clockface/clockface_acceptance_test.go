package clockface_test

import (
	"alnoor/gotdd/clockface"
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  []Circle `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

func TestSVGWriterForSecondsHand(t *testing.T) {
	testCases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("check seconds line for %v to output %v", tC.time, tC.line), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tC.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(svg, tC.line) {
				t.Errorf("Expected to find the second hand with %+v, in the SVG output %v", tC.line, b.String())
			}

		})
	}
}

// func TestSVGWriterForMinutesHand(t *testing.T) {
// 	testCases := []struct {
// 		time time.Time
// 		line Line
// 	}{
// 		{
// 			simpleTime(0, 0, 0),
// 			Line{150, 150, 150, 70},
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(fmt.Sprintf("check seconds line for %v to output %v", tC.time, tC.line), func(t *testing.T) {
// 			b := bytes.Buffer{}
// 			clockface.SVGWriter(&b, tC.time)

// 			svg := SVG{}
// 			xml.Unmarshal(b.Bytes(), &svg)

// 			if !containsLine(svg, tC.line) {
// 				t.Errorf("Expected to find the second hand with %+v, in the SVG output %v", tC.line, b.String())
// 			}

// 		})
// 	}
// }

func containsLine(svg SVG, caseLine Line) bool {
	for _, line := range svg.Line {
		if line == caseLine {
			return true
		}
	}
	return false
}
