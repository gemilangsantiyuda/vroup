package iwd

import "github.com/vroup/preprocess2/domain/model"

//WaterDrop the water drop in the IWD
type WaterDrop struct {
	Soil            float64
	Velocity        float64
	VisitedIndexMap map[int]bool
	RouteList       []model.Route
}
