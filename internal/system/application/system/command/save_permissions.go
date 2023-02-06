package command

type SavePermissions struct {
	RoleId        int64    `json:"roleId,string"`
	PermissionIds []string `json:"permissionIds"`
}
