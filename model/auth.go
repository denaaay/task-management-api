package model

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPayload struct {
	Id   int
	Name string
}

type LoginResp struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
