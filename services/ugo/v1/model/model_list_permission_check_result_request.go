package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPermissionCheckResultRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id" xml:"migration_project_id"`

	// 分页查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListPermissionCheckResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionCheckResultRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionCheckResultRequest", string(data)}, " ")
}
