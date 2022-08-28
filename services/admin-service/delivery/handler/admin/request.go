package admin

type loginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
