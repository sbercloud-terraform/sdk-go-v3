package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPremiumHostResponse struct {

	// 全部防护域名的数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 详细的防护域名信息数组
	Items          *[]SimplePremiumWafHost `json:"items,omitempty" xml:"items"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListPremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPremiumHostResponse struct{}"
	}

	return strings.Join([]string{"ListPremiumHostResponse", string(data)}, " ")
}
