package vo

type UserReq struct {
	Name     string
	Password string
	Nickname string
}

type UserVO struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}
