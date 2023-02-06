package query

type ListSystemRole struct {
	Name        string `json:"name"`        // 角色名称
	OrderBy     string `json:"orderBy"`     // 排序规则
	OrderColumn string `json:"orderColumn"` // 排序字段
	Page        int    `json:"page"`        // 页码
	PageSize    int    `json:"pageSize"`    // 每页行数
}
