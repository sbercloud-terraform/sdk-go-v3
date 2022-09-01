package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronShowFirewallGroupResponse struct {
	FirewallGroup  *NeutronFirewallGroup `json:"firewall_group,omitempty" xml:"firewall_group"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronShowFirewallGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallGroupResponse", string(data)}, " ")
}
