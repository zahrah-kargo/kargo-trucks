// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Shipment struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	DeliveryDate string `json:"delivery_date"`
	TruckID      string `json:"truckId"`
}

type Truck struct {
	ID        string `json:"id"`
	PlateNo   string `json:"plateNo"`
	IsDeleted bool   `json:"isDeleted"`
}
