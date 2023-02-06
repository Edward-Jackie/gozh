package command

// UpdateSettingDictValues 修改
type UpdateSettingDictValues struct {
	// ID
	Id int64 `json:"id,string" label:"ID" validate:"required"`
	// 字典ID
	DictId int64 `json:"dictId,string" label:"字典ID" validate:"required"`
	// 字典值Key
	Key string `json:"key" label:"字典值Key" validate:"required"`
	// 字典值Value
	Value string `json:"value" label:"字典值Value" validate:"required"`
}
