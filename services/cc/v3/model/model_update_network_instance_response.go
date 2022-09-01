package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNetworkInstanceResponse struct {
	NetworkInstance *NetworkInstance `json:"network_instance,omitempty" xml:"network_instance"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNetworkInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateNetworkInstanceResponse", string(data)}, " ")
}
