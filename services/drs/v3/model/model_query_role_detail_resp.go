package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移角色响应体
type QueryRoleDetailResp struct {

	// 角色。
	Role *string `json:"role,omitempty" xml:"role"`

	// 说明。
	Comment *string `json:"comment,omitempty" xml:"comment"`

	// 是否支持迁移。
	IsTransfer *bool `json:"is_transfer,omitempty" xml:"is_transfer"`

	// 角色权限。
	Privileges *string `json:"privileges,omitempty" xml:"privileges"`

	// 继承的角色。
	InheritsRoles *[]string `json:"inherits_roles,omitempty" xml:"inherits_roles"`

	// 是否选择。
	Selected *bool `json:"selected,omitempty" xml:"selected"`
}

func (o QueryRoleDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRoleDetailResp struct{}"
	}

	return strings.Join([]string{"QueryRoleDetailResp", string(data)}, " ")
}
