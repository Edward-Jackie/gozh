package entity

// SettingDict 数据字典
type SettingDict struct {
	// ID
	Id int64 `json:"id,string"`
	// 字典编码
	Code string `json:"code"`
	// 字典名称
	Name string `json:"name"`
}

// Update 更新数据
func (settingDict *SettingDict) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		settingDict.Id = id.(int64)
	}
	// 字典编码
	if code, ok := options["Code"]; ok {
		settingDict.Code = code.(string)
	}
	// 字典名称
	if name, ok := options["Name"]; ok {
		settingDict.Name = name.(string)
	}
}
