package service

// RegisterRequest注册请求
type RegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"pass_word"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	NickName string `json:"nick_name"`
}

// LoginRequest
// @Description: 登录请求
type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"pass_word"`
}
