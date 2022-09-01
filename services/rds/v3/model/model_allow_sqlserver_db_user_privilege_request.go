package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AllowSqlserverDbUserPrivilegeRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *SqlserverGrantRequest `json:"body,omitempty" xml:"body"`
}

func (o AllowSqlserverDbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowSqlserverDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"AllowSqlserverDbUserPrivilegeRequest", string(data)}, " ")
}
