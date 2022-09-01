package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountStatus struct {

	// 是否有创建实例权限
	CurOrgCreateRole *bool `json:"cur_org_create_role,omitempty" xml:"cur_org_create_role"`

	// 帐号所属租户是否开通服务
	CurOrgOpen *bool `json:"cur_org_open,omitempty" xml:"cur_org_open"`

	// 免费试用
	HasFreeTrial *bool `json:"has_free_trial,omitempty" xml:"has_free_trial"`

	// 是否有管理入口的权限
	ShowManage *bool `json:"show_manage,omitempty" xml:"show_manage"`
}

func (o AccountStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountStatus struct{}"
	}

	return strings.Join([]string{"AccountStatus", string(data)}, " ")
}
