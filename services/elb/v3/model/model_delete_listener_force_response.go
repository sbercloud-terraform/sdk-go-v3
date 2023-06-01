package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteListenerForceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerForceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerForceResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerForceResponse", string(data)}, " ")
}
