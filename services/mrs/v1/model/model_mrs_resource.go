package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MrsResource struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源详情
	ResourceDetail *string `json:"resource_detail,omitempty" xml:"resource_detail"`

	// 标签
	Tags *[]TagPlain `json:"tags,omitempty" xml:"tags"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`
}

func (o MrsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrsResource struct{}"
	}

	return strings.Join([]string{"MrsResource", string(data)}, " ")
}
