package entity

import (
	"math/rand"
	"strconv"
)

type Point struct {
	Time     string  `json:"time"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Pressure int     `json:"pressure"`
}

type Stroke struct {
	PenColor             string  `json:"penColor"`
	DotSize              float64 `json:"dotSize"`
	MinWidth             float64 `json:"minWidth"`
	MaxWidth             float64 `json:"maxWidth"`
	VelocityFilterWeight float64 `json:"velocityFilterWeight"`
	CompositeOperation   string  `json:"compositeOperation"`
	Points               []Point `json:"points"`
}

type Strokes []Stroke

// テスト用
func MakePoint() Point {
	return Point{
		Time:     strconv.FormatInt(rand.Int63(), 10),
		X:        rand.Float64(),
		Y:        rand.Float64(),
		Pressure: 1,
	}
}

// テスト用
func MakeStroke() Stroke {
	n := int(rand.Int31n(10))
	points := make([]Point, 0, n)
	for i := 0; i < n; i++ {
		points = append(points, MakePoint())
	}

	return Stroke{
		PenColor:             "black",
		DotSize:              rand.Float64(),
		MinWidth:             rand.NormFloat64(),
		MaxWidth:             rand.NormFloat64(),
		VelocityFilterWeight: rand.NormFloat64(),
		CompositeOperation:   "source-over",
		Points:               points,
	}
}

// テスト用
func MakeStrokes() Strokes {
	n := int(rand.Int31n(10))
	strokes := make([]Stroke, 0, n)
	for i := 0; i < n; i++ {
		strokes = append(strokes, MakeStroke())
	}
	return strokes
}
