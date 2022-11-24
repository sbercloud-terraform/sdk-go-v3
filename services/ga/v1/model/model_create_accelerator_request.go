package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAcceleratorRequest struct {
	Body *CreateAcceleratorRequestBody `json:"body,omitempty"`
}

func (o CreateAcceleratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceleratorRequest struct{}"
	}

	return strings.Join([]string{"CreateAcceleratorRequest", string(data)}, " ")
}
