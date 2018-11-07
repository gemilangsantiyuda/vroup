package model

//Node node struct for the input of the algorithm
type Node struct {
	ID         string
	Coordinate Coordinate
	Qty        int
	Capacity   KitchenCapacity
}
