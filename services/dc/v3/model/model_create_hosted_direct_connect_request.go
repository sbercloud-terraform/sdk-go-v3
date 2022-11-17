package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateHostedDirectConnectRequest struct {
	Body *CreateHostedDirectConnectRequestBody `json:"body,omitempty"`
}

func (o CreateHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnectRequest", string(data)}, " ")
}
