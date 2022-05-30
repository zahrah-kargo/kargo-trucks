package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"kargo-trucks/graph/generated"
	"kargo-trucks/graph/model"
	"strconv"
	"strings"
)

func validatePlateNo(plateNo string) error {
	plateParts := strings.Split(plateNo, " ")

	if len(plateParts) != 3 {
		return errors.New(PLATE_NO_INVALID)
	}

	if len(plateParts) > 2 {
		if len(plateParts[0]) > 2 {
			return errors.New(PLATE_NO_INVALID)
		}

		number, err := strconv.Atoi(plateParts[1])
		if err != nil {
			return errors.New(PLATE_NO_INVALID)
		}

		if number < 0 || number > 9999 {
			return errors.New(PLATE_NO_INVALID)
		}

		if len(plateParts[2]) > 3 {
			return errors.New(PLATE_NO_INVALID)
		}
	}

	return nil

}

func (r *mutationResolver) findTruckByID(truckID string) (*model.Truck, error) {
	if truck, ok := r.Trucks[truckID]; ok && !truck.IsDeleted {
		return &truck, nil
	}
	return nil, errors.New(TRUCK_UNAVAILABLE)

}

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	err := validatePlateNo(plateNo)
	if err != nil {
		return nil, err
	}

	truck := model.Truck{
		ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo: plateNo,
	}
	r.Trucks[truck.ID] = truck
	return &truck, nil

}

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID string) (*model.Shipment, error) {

	_, err := r.findTruckByID(truckID)
	if err != nil {
		return nil, err
	}

	shipment := model.Shipment{
		ID:           fmt.Sprintf("Shipment-%d", len(r.Shipments)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		TruckID:      truckID,
	}
	r.Shipments[shipment.ID] = shipment
	return &shipment, nil

}

func (r *queryResolver) PaginatedTrucks(ctx context.Context, id *string, plateNo *string, page int, first int) ([]*model.Truck, error) {
	trucks := make([]*model.Truck, 0)

	if id != nil || plateNo != nil {
		if truck, ok := r.Trucks[*id]; ok {
			trucks = append(trucks, &truck)
			return trucks, nil
		}
	} else {
		for _, t := range r.Trucks {
			trucks = append(trucks, &t)
		}
		return trucks, nil
	}

	return trucks, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, id *string, origin *string, page int, first int) ([]*model.Shipment, error) {
	shipments := make([]*model.Shipment, 0)

	if id != nil || origin != nil {
		if shipment, ok := r.Shipments[*id]; ok {
			shipments = append(shipments, &shipment)
			return shipments, nil
		}
	} else {
		for _, s := range r.Shipments {
			shipments = append(shipments, &s)
		}
		return shipments, nil
	}

	return shipments, nil
}

func (r *mutationResolver) DeleteTruck(ctx context.Context, id string) (bool, error) {
	if truck, ok := r.Trucks[id]; ok {
		truck.IsDeleted = true
		r.Trucks[id] = truck
		return true, nil
	}
	return false, errors.New(TRUCK_UNAVAILABLE)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
