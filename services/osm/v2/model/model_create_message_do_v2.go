package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMessageDoV2 struct {

	// 留言内容
	Content string `json:"content" xml:"content"`

	// 是否授权
	IsAuthorized *int32 `json:"is_authorized,omitempty" xml:"is_authorized"`

	// 机密信息
	AuthorizationContent *string `json:"authorization_content,omitempty" xml:"authorization_content"`

	// 附件id
	AccessoryIds *[]string `json:"accessory_ids,omitempty" xml:"accessory_ids"`
}

func (o CreateMessageDoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageDoV2 struct{}"
	}

	return strings.Join([]string{"CreateMessageDoV2", string(data)}, " ")
}
