package entity



type User struct {
	UserReference   string      `json:"user_reference" bson:"user_reference"`
	UserFullName    string 	  	`json:"user_full_name" bson:"user_full_name"`
	Unit            string 		`json:"unit" bson:"unit"`
	Role            string      `json:"role" bson:"role"`
	Department      string 		`json:"department" bson:"department"`
	Email 			string 		`json:"email" bson:"email"`
	Phone 			string 		`json:"phone" bson:"phone"`
	IPAddress 		string 		`json:"ip_address" bson:"ip_address"`
    Location        string 	  	`json:"location" bson:"location"`
	ComputerName    string 	    `json:"computer_name" bson:"computer_name"`
}
