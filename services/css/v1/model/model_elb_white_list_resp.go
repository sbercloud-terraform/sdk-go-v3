package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 公网访问信息。
type ElbWhiteListResp struct {

	// 是否开启访问控制。
	EnableWhiteList *bool `json:"enableWhiteList,omitempty"`

	// 公网访问白名单。
	WhiteList *string `json:"whiteList,omitempty"`
}

func (o ElbWhiteListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbWhiteListResp struct{}"
	}

	return strings.Join([]string{"ElbWhiteListResp", string(data)}, " ")
}
