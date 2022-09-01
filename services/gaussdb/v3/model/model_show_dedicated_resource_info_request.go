package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDedicatedResourceInfoRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 专属资源池ID。
	DedicatedResourceId string `json:"dedicated_resource_id" xml:"dedicated_resource_id"`
}

func (o ShowDedicatedResourceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedResourceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowDedicatedResourceInfoRequest", string(data)}, " ")
}
