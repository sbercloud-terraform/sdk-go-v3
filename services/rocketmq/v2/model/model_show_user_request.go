package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 用户名。
	UserName string `json:"user_name" xml:"user_name"`
}

func (o ShowUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRequest struct{}"
	}

	return strings.Join([]string{"ShowUserRequest", string(data)}, " ")
}
