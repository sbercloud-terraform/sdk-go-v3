package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSubscriptionRequest struct {
	Body *SubscriptionCreateReq `json:"body,omitempty" xml:"body"`
}

func (o CreateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionRequest", string(data)}, " ")
}
