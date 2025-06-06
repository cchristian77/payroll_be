package request

type Login struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required"`
}
