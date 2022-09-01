package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResourcesDetailResponseBody struct {

	// 配额资源类型，当前配额类型仅支持实例类型（instance）。
	Type string `json:"type" xml:"type"`

	// 当前配额值。 取值为0时，表示不限制当前配额值。
	Quota int32 `json:"quota" xml:"quota"`

	// 已使用的资源数。
	Used int32 `json:"used" xml:"used"`
}

func (o ShowResourcesDetailResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesDetailResponseBody struct{}"
	}

	return strings.Join([]string{"ShowResourcesDetailResponseBody", string(data)}, " ")
}
