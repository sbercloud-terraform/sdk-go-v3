package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePostPaidInstanceRequest struct {
	Body *CreatePostPaidInstanceReq `json:"body,omitempty" xml:"body"`
}

func (o CreatePostPaidInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceRequest", string(data)}, " ")
}
