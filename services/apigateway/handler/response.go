package handler

//Admin Response
type CreateAdminResponse struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}

type LoginAdminResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type UpdateAdminResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//User Response
type LoginUserResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type GetUserResponse struct {
	IdString string `json:"id" bson:"-"`
	Email    string `bson:"email"`
	Name     string `bson:"name"`
}

type CreateUserResponse struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}
