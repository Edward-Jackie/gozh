package command

type CreateSystemEvent struct {
	// 埋点事件
	Event string `json:"event"`
	// 访问IP
	Ip string `json:"ip"`
	// 关联ID
	RelationId int64 `json:"relationId,string"`
}
