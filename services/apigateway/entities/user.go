package entities

type User struct {
	IdString string `json:"id" bson:"-"`
	Email    string `bson:"email"`
	Name     string `bson:"name"`
}

type UserResponse struct {
	Status string `json:"status"`
	Data   Admin  `json:"data"`
}
