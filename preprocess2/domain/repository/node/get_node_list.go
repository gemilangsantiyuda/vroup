package node

import (
	"strconv"

	"github.com/vroup/preprocess2/domain/model"
)

//GetNodeList transforming orders and kitchens into input nodes for the algorithm
func GetNodeList(orderList []model.Order, kitchenList []model.Kitchen) []model.Node {

	var nodeList []model.Node

	nodeList = addOrderToNodeList(orderList, nodeList)
	nodeList = addKitchenToNodeList(kitchenList, nodeList)

	return nodeList
}

func addOrderToNodeList(orderList []model.Order, nodeList []model.Node) []model.Node {

	for _, order := range orderList {

		id := strconv.Itoa(order.ID)
		coordinate := order.Coordinate
		qty := order.Qty
		capacity := model.KitchenCapacity{
			Minimum: 0,
			Optimum: 0,
			Maximum: 0,
		}

		node := model.Node{
			ID:         id,
			Coordinate: coordinate,
			Qty:        qty,
			Capacity:   capacity,
		}

		/* 	Jika customer memesan lebih dari 40 meal
		maka ia akan dipisah menjadi beberapa node
		yang masing-masing memesan <=40 meal
		pemecahan dilakukan dengan membuat node
		baru yang memesan 40 meal, dan mengurangi
		pesanan dari node aslinya sejumlah 40 meal
		hingga node aslinya memesan <= 40 meal
		*/
		if node.Qty > 40 {
			nodeCopy := node
			nodeCopy.Qty = 40
			for node.Qty > 40 {
				nodeList = append(nodeList, nodeCopy)
				node.Qty -= 40
			}
		}

		nodeList = append(nodeList, node)
	}

	return nodeList
}

func addKitchenToNodeList(kitchenList []model.Kitchen, nodeList []model.Node) []model.Node {

	for _, kitchen := range kitchenList {

		id := strconv.Itoa(kitchen.ID)
		coordinate := kitchen.Coordinate
		capacity := kitchen.Capacity

		node := model.Node{
			ID:         id,
			Coordinate: coordinate,
			Qty:        0,
			Capacity:   capacity,
		}

		nodeList = append(nodeList, node)
	}

	return nodeList
}
