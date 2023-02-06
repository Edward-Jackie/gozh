package command

// UpdateSettingDict 修改数据字典
type UpdateSettingDict struct {
	// ID
	Id int64 `json:"id,string" label:"ID"`
	// 字典编码
	Code string `json:"code" label:"字典编码"`
	// 字典名称
	Name string `json:"name" label:"字典名称"`
}
