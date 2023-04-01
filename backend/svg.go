package main

import (
	"bytes"
	"encoding/xml"

	svg "github.com/ajstarks/svgo"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Polygon string   `xml:"polygon"`
	Rects   []string `xml:"rect"`
}

func createSVG(polygon []Point, rectangles []Rectangle) ([]byte, error) {
	_, _, width, height := PoligonBB(polygon)
	// width := int(polygon[2].X)
	// height := int(polygon[2].Y)

	var buffer bytes.Buffer
	canvas := svg.New(&buffer)
	canvas.Start(int(width), int(height))
	canvas.Gstyle("fill:#FFA500;stroke:black")

	// Draw polygon
	canvas.Polygon(polygonToSVGPoints(polygon))

	// Draw rectangles
	for _, r := range rectangles {
		canvas.Rect(int(r.TopLeft.X), int(r.TopLeft.Y), int(r.Width), int(r.Height), "fill:yellow;stroke:black")
	}

	canvas.Gend()
	canvas.End()

	return buffer.Bytes(), nil
}

func polygonToSVGPoints(polygon []Point) (x, y []int) {
	x = make([]int, len(polygon))
	y = make([]int, len(polygon))
	for i, p := range polygon {
		x[i] = int(p.X)
		y[i] = int(p.Y)
	}
	return
}
