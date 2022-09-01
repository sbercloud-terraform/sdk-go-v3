package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 已授权数据库用户信息。
type GaussMySqlDatabaseInfo struct {

	// 数据库用户名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 主机地址。
	Host *string `json:"host,omitempty" xml:"host"`

	// 是否为只读权限：true表示只读,false表示可读写。
	Readonly *bool `json:"readonly,omitempty" xml:"readonly"`
}

func (o GaussMySqlDatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussMySqlDatabaseInfo struct{}"
	}

	return strings.Join([]string{"GaussMySqlDatabaseInfo", string(data)}, " ")
}
