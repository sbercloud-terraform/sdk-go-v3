package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResizeInstanceResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 订单ID，仅变更包年包月实例的规格时返回该参数。
	OrderId        *string `json:"order_id,omitempty" xml:"order_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceResponse", string(data)}, " ")
}
