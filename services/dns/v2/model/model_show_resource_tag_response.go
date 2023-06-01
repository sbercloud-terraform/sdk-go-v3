package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResourceTagResponse struct {

	// 指定实例的标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目或默认项目
	EnterpriseProjectOrDefault *string `json:"enterpriseProjectOrDefault,omitempty"`
	HttpStatusCode             int     `json:"-"`
}

func (o ShowResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagResponse", string(data)}, " ")
}
