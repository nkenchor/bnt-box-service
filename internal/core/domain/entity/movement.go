package entity

type Movement struct {
	Reference       string     	`json:"reference" bson:"reference"`
	MovementBatchNo	string		`json:"movement_batch_no" bson:"movement_batch_no"`
	VehicleType		string		`json:"vehicle_type" bson:"vehicle_type"`
	VehicleNo		string		`json:"vehicle_no" bson:"vehicle_no"`
	DriverLastName  string		`json:"driver_last_name" bson:"driver_last_name"`
	DriverFirstName string		`json:"driver_first_name" bson:"driver_first_name"`
	CompanyName 	string		`json:"company_name" bson:"company_name"`
	Destination		string		`json:"destination" bson:"destination"`
	Documentation   Document	`json:"documentation" bson:"documentation"`
}