package entities

type Admin struct {
	IdString string `json:"id" bson:"-"`
	Email    string `bson:"email"`
	Name     string `bson:"name"`
}

type AdminResponse struct {
	Status string `json:"status"`
	Data   Admin  `json:"data"`
}

type UpdateAdminRequest struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
