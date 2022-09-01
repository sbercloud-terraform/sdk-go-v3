package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPoliciesResponse struct {

	//
	Policies *[]Policy `json:"policies,omitempty" xml:"policies"`

	// 策略总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListPoliciesResponse", string(data)}, " ")
}
