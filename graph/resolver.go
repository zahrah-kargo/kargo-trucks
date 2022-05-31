package graph

import (
	"fmt"
	"kargo-trucks/graph/model"
	"math/rand"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Trucks    map[string]model.Truck
	Shipments map[string]model.Shipment
}

func (r *Resolver) Init() {
	for i := 0; i < 20; i++ {
		num := rand.Intn(100)
		truck := model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", num),
			PlateNo: fmt.Sprintf("B %d CD", num),
		}
		r.Trucks[truck.ID] = truck
	}
}
