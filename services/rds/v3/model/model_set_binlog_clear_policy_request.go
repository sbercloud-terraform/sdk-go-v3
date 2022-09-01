package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetBinlogClearPolicyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *BinlogClearPolicyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o SetBinlogClearPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBinlogClearPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetBinlogClearPolicyRequest", string(data)}, " ")
}
