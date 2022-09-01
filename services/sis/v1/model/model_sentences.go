package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Sentences struct {

	// 分句结果信息
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time"`

	Result *FlashScoreResult `json:"result,omitempty" xml:"result"`

	// 一句话开始时间，单位毫秒
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time"`
}

func (o Sentences) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sentences struct{}"
	}

	return strings.Join([]string{"Sentences", string(data)}, " ")
}
