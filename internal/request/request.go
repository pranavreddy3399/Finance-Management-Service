package request

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phno  string `json:"phno"`
}
