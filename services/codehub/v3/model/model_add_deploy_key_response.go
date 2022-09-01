package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddDeployKeyResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *Key `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDeployKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyResponse struct{}"
	}

	return strings.Join([]string{"AddDeployKeyResponse", string(data)}, " ")
}
