package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDeployKeyV2Response struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 响应结果
	Result *bool `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeployKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployKeyV2Response struct{}"
	}

	return strings.Join([]string{"DeleteDeployKeyV2Response", string(data)}, " ")
}
