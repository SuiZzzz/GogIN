package types

type AIReq struct {
	Data   string `json:"data" form:"data"`
	UserId string `json:"user_id" form:"user_id"`
}
