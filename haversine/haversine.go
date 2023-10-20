package haversine

import "math"

type Pairs struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	X0 float64 `json:"x0"`
	X1 float64 `json:"x1"`
	Y0 float64 `json:"y0"`
	Y1 float64 `json:"y1"`
}

func square(a float64) (s float64) {
	s = a * a
	return s
}

func radiansFromDegrees(degrees float64) (radians float64) {
	radians = 0.01745329251994329577 * degrees
	return
}

func ReferenceHaversine(x0, x1, y0, y1, earthRadius float64) (distance float64) {
	var lat1, lat2, lon1, lon2 float64 = y0, y1, x0, x1

	var dLat float64 = radiansFromDegrees(lat2 - lat1)
	var dLon float64 = radiansFromDegrees(lon2 - lon1)
	lat1 = radiansFromDegrees(lat1)
	lat2 = radiansFromDegrees(lat2)

	var a float64 = square(math.Sin(dLat/2.0)) + math.Cos(lat1)*math.Cos(lat2)*square(math.Sin(dLon/2.0))
	var c float64 = 2.0 * math.Asin(math.Sqrt(a))

	distance = c * earthRadius
	return
}
