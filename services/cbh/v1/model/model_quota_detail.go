package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 当前租户CBH的配额信息
type QuotaDetail struct {

	// 中文配额描述
	ZhCn *string `json:"zh_cn,omitempty" xml:"zh_cn"`

	// 英文配额描述
	EnUs *string `json:"en_us,omitempty" xml:"en_us"`

	// 租户剩余配额数量
	Remaining int32 `json:"remaining" xml:"remaining"`
}

func (o QuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
