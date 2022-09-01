package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateCertificateV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id" xml:"group_id"`

	// 域名的编号
	DomainId string `json:"domain_id" xml:"domain_id"`

	Body *CertForm `json:"body,omitempty" xml:"body"`
}

func (o AssociateCertificateV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateCertificateV2Request struct{}"
	}

	return strings.Join([]string{"AssociateCertificateV2Request", string(data)}, " ")
}
