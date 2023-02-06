package entity

// SettingDictValues 字典值
type SettingDictValues struct {
	// ID
	Id int64 `json:"id,string"`
	// 字典ID
	DictId int64 `json:"dictId,string"`
	// 字典值Key
	Key string `json:"key"`
	// 字典值Value
	Value string `json:"value"`
}

type DictOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Update 更新数据
func (settingDictValues *SettingDictValues) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		settingDictValues.Id = id.(int64)
	}
	// 字典ID
	if dictId, ok := options["DictId"]; ok {
		settingDictValues.DictId = dictId.(int64)
	}
	// 字典值Key
	if key, ok := options["Key"]; ok {
		settingDictValues.Key = key.(string)
	}
	// 字典值Value
	if value, ok := options["Value"]; ok {
		settingDictValues.Value = value.(string)
	}
}
