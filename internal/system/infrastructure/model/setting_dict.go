package model

// SettingDict 数据字典
type SettingDict struct {
	// ID
	Id int64 `json:"id"`
	// 字典编码
	Code string `json:"code"`
	// 字典名称
	Name string `json:"name"`
}
