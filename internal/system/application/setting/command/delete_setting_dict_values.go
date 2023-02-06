package command

// DeleteSettingDictValues 删除
type DeleteSettingDictValues struct {
	// ID
	Id  int64    `json:"id,string" label:"ID"`
	Ids []string `json:"ids" label:"Ids"`
}
