package handler

//Admin Request
type CreateAdminRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
}

type LoginAdminRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UpdateAdminRequest struct {
	Name     string `json:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

//User Request
type LoginUserRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type CreateUserRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
}
