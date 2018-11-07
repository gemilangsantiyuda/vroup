package distance

import (
	"math"

	"github.com/vroup/preprocess2/domain/model"
)

//CalculateHaversineDistance just calculate haversine distance of 2 given coordinates
func CalculateHaversineDistance(coordOrigin, coordDestination model.Coordinate) float64 {
	DY := math.Abs(coordOrigin.Latitude-coordDestination.Latitude) / 180 * math.Pi
	DX := math.Abs(coordOrigin.Longitude-coordDestination.Longitude) / 180 * math.Pi
	Y1 := coordOrigin.Latitude / 180 * math.Pi
	Y2 := coordDestination.Latitude / 180 * math.Pi
	R := 6372800.00000000 //in Meter ! it is the average great-elliptic or great-circle radius
	a := math.Sin(DY/2)*math.Sin(DY/2) + math.Cos(Y1)*math.Cos(Y2)*math.Sin(DX/2)*math.Sin(DX/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
