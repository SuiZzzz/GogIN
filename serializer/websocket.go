package serializer

const (
	Notification = 1 << iota
	Group
)

type ClientReq struct {
	UserId uint `json:"user_id"`
	Type   uint `json:"type"`
}

type ClientResp struct {
	Audit    []string `json:"audit,omitempty"`
	Message  []string `json:"message,omitempty"`
	Code     uint     `json:"code,omitempty"`
	Error    string   `json:"error,omitempty"`
	Nickname []string `json:"Nickname,omitempty"`
}

//type Member struct {
//	Nickname string `json:"nickname,omitempty"`
//	Level    string `json:"level,omitempty"`
//}
