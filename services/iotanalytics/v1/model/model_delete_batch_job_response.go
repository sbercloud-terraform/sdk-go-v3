package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteBatchJobResponse struct {

	// 被删除作业ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBatchJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteBatchJobResponse", string(data)}, " ")
}
