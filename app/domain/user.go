package domain

type DemoTestReq struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type DemoTestRsp struct {
	Msg string `json:"msg"`
}

type RegisterOrUpdateReq struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginReq struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
}

type LoginRsp struct {
	SessionId string `json:"sessionId"`
}

type GetUserInfoReq struct {
	Passport string `json:"passport"`
}

type GetUserInfoRsp struct {
	Passport string `json:"passport"`
	Password uint64 `json:"password"`
	Nickname string `json:"nickname"`
}
