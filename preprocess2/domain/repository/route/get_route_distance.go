package route

import (
	"github.com/vroup/preprocess2/domain/model"
	distanceRepo "github.com/vroup/preprocess2/domain/repository/distance"
)

//GetRouteDistance Calculate total distance travelled in a route
func GetRouteDistance(route model.Route, distanceMap map[string]float64) float64 {

	distance := 0.0
	for idx := 1; idx < len(route); idx++ {
		coordSource := route[idx-1].Coordinate
		coordDestination := route[idx].Coordinate
		dist := distanceRepo.GetHaversineDistance(coordSource, coordDestination, distanceMap)
		distance += dist
	}
	return distance
}
