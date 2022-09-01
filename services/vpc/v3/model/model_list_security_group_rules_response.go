package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecurityGroupRulesResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 安全组规则列表响应体
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules,omitempty" xml:"security_group_rules"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesResponse", string(data)}, " ")
}
