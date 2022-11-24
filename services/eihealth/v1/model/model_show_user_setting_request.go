package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowUserSettingRequest struct {

	// 用户id
	UserId string `json:"user_id"`
}

func (o ShowUserSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowUserSettingRequest", string(data)}, " ")
}
