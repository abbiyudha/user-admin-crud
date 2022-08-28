package admin

type GetAdminResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateAdminResponse struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}

type UnAuthorizeResponse struct {
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
