package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostResponse struct {

	// 云模式防护域名的数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 详细的云模式防护域名列表信息
	Items          *[]CloudWafHostItem `json:"items,omitempty" xml:"items"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostResponse struct{}"
	}

	return strings.Join([]string{"ListHostResponse", string(data)}, " ")
}
