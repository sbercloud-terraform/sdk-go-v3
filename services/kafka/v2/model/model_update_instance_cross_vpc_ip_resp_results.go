package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改broker跨VPC访问的结果。
type UpdateInstanceCrossVpcIpRespResults struct {

	// advertised.listeners IP/域名。
	AdvertisedIp *string `json:"advertised_ip,omitempty" xml:"advertised_ip"`

	// 修改broker跨VPC访问的状态。
	Success *bool `json:"success,omitempty" xml:"success"`

	// listeners IP。
	Ip *string `json:"ip,omitempty" xml:"ip"`
}

func (o UpdateInstanceCrossVpcIpRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpRespResults struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpRespResults", string(data)}, " ")
}
