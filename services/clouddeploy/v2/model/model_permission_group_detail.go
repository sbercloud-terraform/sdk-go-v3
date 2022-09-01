package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主机组相关权限详情类
type PermissionGroupDetail struct {

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty" xml:"can_view"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty" xml:"can_edit"`

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty" xml:"can_delete"`

	// 是否有添加主机权限
	CanAddHost *bool `json:"can_add_host,omitempty" xml:"can_add_host"`

	// 是否有管理权限
	CanManage *bool `json:"can_manage,omitempty" xml:"can_manage"`
}

func (o PermissionGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionGroupDetail struct{}"
	}

	return strings.Join([]string{"PermissionGroupDetail", string(data)}, " ")
}
