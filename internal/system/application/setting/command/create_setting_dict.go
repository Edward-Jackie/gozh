package command

// CreateSettingDict 新增数据字典
type CreateSettingDict struct {
	// 字典编码
	Code string `json:"code" label:"字典编码"`
	// 字典名称
	Name string `json:"name" label:"字典名称"`
}
