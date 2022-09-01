package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新容灾演练名称数据结构
type UpdateDisasterRecoveryDrillNameRequestParams struct {

	// 容灾演练的名称。最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	Name string `json:"name" xml:"name"`
}

func (o UpdateDisasterRecoveryDrillNameRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryDrillNameRequestParams struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryDrillNameRequestParams", string(data)}, " ")
}
