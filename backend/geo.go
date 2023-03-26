package main

import "math"

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	TopLeft Point
	Width   float64
	Height  float64
}

func polygonArea(polygon []Point) float64 {
	area := 0.0
	n := len(polygon)

	for i := 0; i < n-1; i++ {
		area += (polygon[i].X * polygon[i+1].Y) - (polygon[i+1].X * polygon[i].Y)
	}
	area += (polygon[n-1].X * polygon[0].Y) - (polygon[0].X * polygon[n-1].Y)

	return math.Abs(area) / 2
}

func isPointInsidePolygon(polygon []Point, point Point) bool {
	isInside := false
	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {
		if ((polygon[i].Y > point.Y) != (polygon[j].Y > point.Y)) &&
			(point.X < (polygon[j].X-polygon[i].X)*(point.Y-polygon[i].Y)/(polygon[j].Y-polygon[i].Y)+polygon[i].X) {
			isInside = !isInside
		}
		j = i
	}

	return isInside
}

func PoligonBB(polygon []Point) (minX, minY, maxX, maxY float64) {
	minX, minY = math.MaxFloat64, math.MaxFloat64
	maxX, maxY = -math.MaxFloat64, -math.MaxFloat64

	for _, p := range polygon {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return
}

func RectanglesInBoundingBox(bbMin, bbMax Point, rectSize Rectangle) []Point {
	var topLeftPoints []Point

	// Calculate how many rectangles fit horizontally and vertically
	horizontalCount := int((bbMax.X - bbMin.X) / rectSize.Width)
	verticalCount := int((bbMax.Y - bbMin.Y) / rectSize.Height)

	// Iterate through the positions where the rectangles can be placed
	for i := 0; i < verticalCount; i++ {
		for j := 0; j < horizontalCount; j++ {
			topLeft := Point{
				X: bbMin.X + float64(j)*rectSize.Width,
				Y: bbMin.Y + float64(i)*rectSize.Height,
			}
			topLeftPoints = append(topLeftPoints, topLeft)
		}
	}

	return topLeftPoints
}

func rectanglesInsidePolygon(polygon []Point, rectSize Rectangle) []Rectangle {
	var rectangles []Rectangle

	// Find the bounding box of the polygon
	bbMin := polygon[0]
	bbMax := polygon[0]

	for _, vertex := range polygon {
		if vertex.X < bbMin.X {
			bbMin.X = vertex.X
		}
		if vertex.X > bbMax.X {
			bbMax.X = vertex.X
		}
		if vertex.Y < bbMin.Y {
			bbMin.Y = vertex.Y
		}
		if vertex.Y > bbMax.Y {
			bbMax.Y = vertex.Y
		}
	}

	// Calculate how many rectangles fit horizontally and vertically
	horizontalCount := int((bbMax.X - bbMin.X) / rectSize.Width)
	verticalCount := int((bbMax.Y - bbMin.Y) / rectSize.Height)

	// Iterate through the positions where the rectangles can be placed
	for i := 0; i < verticalCount; i++ {
		for j := 0; j < horizontalCount; j++ {
			topLeft := Point{
				X: bbMin.X + float64(j)*rectSize.Width,
				Y: bbMin.Y + float64(i)*rectSize.Height,
			}

			// Check if all four corners of the rectangle are inside the polygon
			if isPointInsidePolygon(polygon, topLeft) &&
				isPointInsidePolygon(polygon, Point{topLeft.X + rectSize.Width, topLeft.Y}) &&
				isPointInsidePolygon(polygon, Point{topLeft.X, topLeft.Y + rectSize.Height}) &&
				isPointInsidePolygon(polygon, Point{topLeft.X + rectSize.Width, topLeft.Y + rectSize.Height}) {
				rect := Rectangle{
					TopLeft: topLeft,
					Width:   rectSize.Width,
					Height:  rectSize.Height,
				}
				rectangles = append(rectangles, rect)
			}
		}
	}

	return rectangles
}

func scaleCoordinates(points []Point, targetWidth, targetHeight float64) []Point {
	// Find the bounding box
	minX, minY := points[0].X, points[0].Y
	maxX, maxY := points[0].X, points[0].Y

	for _, point := range points {
		if point.X < minX {
			minX = point.X
		}
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y < minY {
			minY = point.Y
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	width := maxX - minX
	height := maxY - minY

	scaleX := targetWidth / width
	scaleY := targetHeight / height

	scaledPoints := make([]Point, len(points))

	for i, point := range points {
		scaledPoints[i] = Point{
			X: (point.X - minX) * scaleX,
			Y: (point.Y - minY) * scaleY,
		}
	}

	return scaledPoints
}

func scaleRectangle(rect Rectangle, min, max Point, targetWidth, targetHeight float64) Rectangle {
	width := max.X - min.X
	height := max.Y - min.Y

	scaleX := targetWidth / width
	scaleY := targetHeight / height

	scaledRect := Rectangle{
		TopLeft: Point{
			X: (rect.TopLeft.X - min.X) * scaleX,
			Y: (rect.TopLeft.Y - min.Y) * scaleY,
		},
		Width:  rect.Width * scaleX,
		Height: rect.Height * scaleY,
	}

	return scaledRect
}
