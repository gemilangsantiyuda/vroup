package distance

import (
	"strconv"

	"github.com/vroup/preprocess2/domain/model"
)

//GetHaversineDistance Get the haversine distance from the HVSDistanceMap
//the key is coordinateXcoordinate, if the key doesn't exist calculate the haversine
//distance
func GetHaversineDistance(coordOrigin, coordDestination model.Coordinate, distanceMap map[string]float64) float64 {

	lat1 := strconv.FormatFloat(coordOrigin.Latitude, 'f', -1, 64)
	lng1 := strconv.FormatFloat(coordOrigin.Longitude, 'f', -1, 64)
	coord1 := lat1 + "#" + lng1

	lat2 := strconv.FormatFloat(coordDestination.Latitude, 'f', -1, 64)
	lng2 := strconv.FormatFloat(coordDestination.Longitude, 'f', -1, 64)
	coord2 := lat2 + "#" + lng2

	if coord1 < coord2 {
		tmp := coord1
		coord1 = coord2
		coord2 = tmp
	}

	key := coord1 + "$" + coord2

	if val, ok := distanceMap[key]; ok {
		return val
	}

	distance := CalculateHaversineDistance(coordOrigin, coordDestination)
	distanceMap[key] = distance
	return distance
}
