package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTranscodingsTemplateRequest struct {

	// 推流域名
	Domain string `json:"domain" xml:"domain"`

	// 应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 分页编号，默认为0。
	Page *int32 `json:"page,omitempty" xml:"page"`

	// 每页记录数。  取值范围：1-100。  默认为10。
	Size *int32 `json:"size,omitempty" xml:"size"`
}

func (o ShowTranscodingsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTranscodingsTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTranscodingsTemplateRequest", string(data)}, " ")
}
