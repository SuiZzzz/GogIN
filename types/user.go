package types

type UserRegisterReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Nickname string `json:"nickname" form:"nickname"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
}

type UserLoginReq struct {
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	Phone        string `json:"phone,omitempty" form:"phone"`
	VerifiedCode string `json:"verifiedCode,omitempty" form:"email"`
}
