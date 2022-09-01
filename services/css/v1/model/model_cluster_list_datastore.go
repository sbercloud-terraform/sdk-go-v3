package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据搜索引擎类型。
type ClusterListDatastore struct {

	// 引擎类型，目前只支持elasticsearch。
	Type *string `json:"type,omitempty" xml:"type"`

	// CSS集群引擎版本号。详细请参考CSS[支持的集群版本](css_03_0056.xml)。
	Version *string `json:"version,omitempty" xml:"version"`
}

func (o ClusterListDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterListDatastore struct{}"
	}

	return strings.Join([]string{"ClusterListDatastore", string(data)}, " ")
}
