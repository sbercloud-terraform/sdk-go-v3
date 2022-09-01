package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronShowSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty" xml:"security_group"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronShowSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupResponse", string(data)}, " ")
}
