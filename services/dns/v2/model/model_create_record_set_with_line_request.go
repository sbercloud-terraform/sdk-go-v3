package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRecordSetWithLineRequest struct {

	// 所属zone的ID。
	ZoneId string `json:"zone_id" xml:"zone_id"`

	Body *CreateRecordSetWithLineReq `json:"body,omitempty" xml:"body"`
}

func (o CreateRecordSetWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineRequest", string(data)}, " ")
}
