package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppsRequest struct {

	// 应用id
	BusinessId int64 `json:"business_id" xml:"business_id"`

	// 应用id
	XBusinessId int64 `json:"x-business-id" xml:"x-business-id"`
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}
