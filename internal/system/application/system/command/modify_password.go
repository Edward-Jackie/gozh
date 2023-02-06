package command

type ModifyPassword struct {
	//旧密码
	OldPassword string `json:"oldPassword"`
	//密码
	Password string `json:"password"`
	//确认密码
	ConfirmPassword string `json:"confirmPassword"`
}
