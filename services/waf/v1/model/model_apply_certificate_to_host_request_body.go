package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 绑定证书和域名的请求体，云模式HTTPS域名id（cloud_host_ids）与独享模式HTTPS域名id（premium_host_ids）至少填一项
type ApplyCertificateToHostRequestBody struct {

	// 云模式HTTPS域名id，通过查询云模式防护域名列表（ListHost）接口获取
	CloudHostIds *[]string `json:"cloud_host_ids,omitempty" xml:"cloud_host_ids"`

	// 独享模式HTTPS域名id，通过独享模式域名列表（ListPremiumHost）接口获取
	PremiumHostIds *[]string `json:"premium_host_ids,omitempty" xml:"premium_host_ids"`
}

func (o ApplyCertificateToHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostRequestBody", string(data)}, " ")
}
