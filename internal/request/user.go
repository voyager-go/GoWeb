package request

type UserSignInReq struct {
	Phone    string `json:"phone" validate:"required,numeric,len=11"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

type UserCreateReq struct {
	Phone    string `json:"phone" validate:"required,numeric,len=11"`  // 手机号
	Status   uint8  `json:"status"`                                    // 状态 1-启用 2-禁用
	Nickname string `json:"nickname" validate:"required,min=2,max=60"` // 昵称
	Email    string `json:"email" validate:"required,email"`           // 邮箱
	Password string `json:"password" validate:"required,min=4,max=20"` // 密码
}
