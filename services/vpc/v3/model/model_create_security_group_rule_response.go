package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSecurityGroupRuleResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty" xml:"security_group_rule"`
	HttpStatusCode    int                `json:"-"`
}

func (o CreateSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleResponse", string(data)}, " ")
}
