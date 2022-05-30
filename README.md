# query graphQL

```
query SHIPMENTS {
  paginatedShipments(page: 1, first: 10) {
    id
    name
    origin
    destination
    delivery_date
    truckId
  }
}

query TRUCKS {
  paginatedTrucks(id: "TRUCK-2", page: 1, first: 10) {
    id
    plateNo
  }
}

mutation saveTruck {
  saveTruck(plateNo: "KH 99 GOD") {
    id
    plateNo
  }
}

mutation SAVESHIPMENT {
  saveShipment(
    name: "Zahrah"
    origin: "Jakarta"
    destination: "Surabaya"
    delivery_date: "20220530"
    truckId: "TRUCK-1"
  ) {
    id
    name
    origin
    destination
    delivery_date
    truckId
  }
}

mutation DEL_TRUCK {
  deleteTruck(id: "TRUCK-1")
}
```
