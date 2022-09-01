package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnbindApiForAcl struct {

	// API的认证方式
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type"`

	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty" xml:"run_env_name"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 发布记录的编号
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`

	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// API名称
	Name *string `json:"name,omitempty" xml:"name"`

	// API描述
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty" xml:"run_env_id"`

	// API编号
	Id *string `json:"id,omitempty" xml:"id"`

	// API的请求地址
	ReqUri *string `json:"req_uri,omitempty" xml:"req_uri"`

	// API类型
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 绑定的其他同类型的ACL策略名称
	AclName *string `json:"acl_name,omitempty" xml:"acl_name"`
}

func (o UnbindApiForAcl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindApiForAcl struct{}"
	}

	return strings.Join([]string{"UnbindApiForAcl", string(data)}, " ")
}
