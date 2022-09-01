package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddEipV2Response struct {

	// 弹性公网IP编号
	EipId *string `json:"eip_id,omitempty" xml:"eip_id"`

	// 弹性公网IP
	EipAddress *string `json:"eip_address,omitempty" xml:"eip_address"`

	// 弹性公网IP状态
	EipStatus *string `json:"eip_status,omitempty" xml:"eip_status"`

	// 弹性公网IP(IPV6)
	EipIpv6Address *string `json:"eip_ipv6_address,omitempty" xml:"eip_ipv6_address"`
	HttpStatusCode int     `json:"-"`
}

func (o AddEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipV2Response struct{}"
	}

	return strings.Join([]string{"AddEipV2Response", string(data)}, " ")
}
