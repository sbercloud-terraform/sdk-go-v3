package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSumTableRequest struct {

	// 应用id
	XBusinessId int64 `json:"x-business-id" xml:"x-business-id"`

	Body *SumTableParam `json:"body,omitempty" xml:"body"`
}

func (o ShowSumTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSumTableRequest struct{}"
	}

	return strings.Join([]string{"ShowSumTableRequest", string(data)}, " ")
}
