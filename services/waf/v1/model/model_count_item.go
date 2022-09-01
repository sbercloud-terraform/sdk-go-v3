package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 攻击事件统计结果
type CountItem struct {

	// 类型，包括请求总量（ACCESS）、Bot攻击防护（CRAWLER）、攻击总量（TOTAL_ATTACK）、Web基础防护（WEB_ATTACK）、精准防护（PRECISE）以及CC攻击防护（CC）
	Key *string `json:"key,omitempty" xml:"key"`

	// 数量
	Num *int32 `json:"num,omitempty" xml:"num"`
}

func (o CountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountItem struct{}"
	}

	return strings.Join([]string{"CountItem", string(data)}, " ")
}
