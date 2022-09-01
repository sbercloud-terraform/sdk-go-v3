package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartupInstanceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o StartupInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartupInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartupInstanceResponse", string(data)}, " ")
}
