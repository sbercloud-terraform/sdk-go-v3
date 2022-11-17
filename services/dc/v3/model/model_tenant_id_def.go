package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例所属项目ID。
type TenantIdDef struct {
}

func (o TenantIdDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantIdDef struct{}"
	}

	return strings.Join([]string{"TenantIdDef", string(data)}, " ")
}
