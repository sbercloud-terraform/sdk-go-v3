package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSnapshotDataResponse struct {

	// 采样数据列表。
	SnapshotList *[]SnapshotData `json:"snapshot_list,omitempty" xml:"snapshot_list"`

	// 指定时间区间内截图数量总和。
	Total *int64 `json:"total,omitempty" xml:"total"`

	XRequestId     *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSnapshotDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotDataResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotDataResponse", string(data)}, " ")
}
