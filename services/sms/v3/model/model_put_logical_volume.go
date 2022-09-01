package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改的逻辑卷信息
type PutLogicalVolume struct {

	// 逻辑卷id
	Id string `json:"id" xml:"id"`

	// 是否迁移
	NeedMigration *bool `json:"need_migration,omitempty" xml:"need_migration"`

	// 调整大小
	AdjustSize *int64 `json:"adjust_size,omitempty" xml:"adjust_size"`
}

func (o PutLogicalVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutLogicalVolume struct{}"
	}

	return strings.Join([]string{"PutLogicalVolume", string(data)}, " ")
}
