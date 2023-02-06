package query

type ListSystemUser struct {
	NickName      string `json:"nickName"`      // 昵称
	UserName      string `json:"userName"`      // 名称
	LastLoginDate string `json:"lastLoginDate"` // 登录日期
	OrderBy       string `json:"orderBy"`       // 排序规则
	OrderColumn   string `json:"orderColumn"`   // 排序字段
	Page          int    `json:"page"`          // 页码
	PageSize      int    `json:"pageSize"`      // 每页行数
}
