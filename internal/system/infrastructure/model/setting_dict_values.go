package model

// SettingDictValues 字典值
type SettingDictValues struct {
	// ID
	Id int64 `json:"id"`
	// 字典ID
	DictId int64 `json:"dictId,string"`
	// 字典值Key
	Key string `json:"key"`
	// 字典值Value
	Value string `json:"value"`
}
