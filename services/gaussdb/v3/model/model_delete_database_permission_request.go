package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDatabasePermissionRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabasePermissionRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabasePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabasePermissionRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabasePermissionRequest", string(data)}, " ")
}
