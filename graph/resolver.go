package graph

import "kargo-trucks/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Trucks    map[string]model.Truck
	Shipments map[string]model.Shipment
}
