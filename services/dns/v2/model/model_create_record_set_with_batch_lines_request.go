package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRecordSetWithBatchLinesRequest struct {
	ZoneId string `json:"zone_id" xml:"zone_id"`

	Body *CreateRSetBatchLinesReq `json:"body,omitempty" xml:"body"`
}

func (o CreateRecordSetWithBatchLinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithBatchLinesRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithBatchLinesRequest", string(data)}, " ")
}
