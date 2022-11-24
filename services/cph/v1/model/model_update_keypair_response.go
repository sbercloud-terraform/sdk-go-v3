package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateKeypairResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id"`

	// 任务信息
	Jobs           []interface{} `json:"jobs"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeypairResponse", string(data)}, " ")
}
