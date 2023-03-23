package request

type AdminCreateReq struct {
	Nickname string `json:"nickname"   validate:"required,min=2,max=60"`
	Password string `json:"password"  validate:"required,min=4,max=20"`
	Email    string `json:"email" validate:"required,email"`
}

type AdminSignInReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}
