package command

type ResetPassword struct {
	//用户ID
	UserId int64 `json:"userId,string"`
	//密码
	Password string `json:"password"`
	//确认密码
	ConfirmPassword string `json:"confirmPassword"`
}
